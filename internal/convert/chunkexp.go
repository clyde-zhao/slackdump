package convert

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime/trace"
	"sync"

	"github.com/rusq/fsadapter"
	"github.com/rusq/slack"

	"github.com/rusq/slackdump/v3/internal/chunk"
	"github.com/rusq/slackdump/v3/internal/chunk/transform"
	"github.com/rusq/slackdump/v3/internal/chunk/transform/fileproc"
)

const (
	defWorkers = 8 // default number of goroutines to process channels
)

var (
	ErrEmptyChunk    = errors.New("missing chunk")
	ErrNoLocFunction = errors.New("missing location function")
)

// ChunkToExport is a converter between Chunk and Export formats.  Zero value
// is not usable.
type ChunkToExport struct {
	// src is the source directory with chunks
	src *chunk.Directory
	// trg is the target FS for the export
	trg fsadapter.FS
	// UploadDir is the upload directory name (relative to Src)
	includeFiles bool
	// srcFileLoc should return the file location within the source directory.
	srcFileLoc func(*slack.Channel, *slack.File) string
	// trgFileLoc should return the file location within the target directory
	trgFileLoc func(*slack.Channel, *slack.File) string

	lg *slog.Logger

	workers int // number of workers to use to convert channels

	request chan copyrequest
	result  chan copyresult
}

type C2EOption func(*ChunkToExport)

// WithIncludeFiles sets the IncludeFiles option.
func WithIncludeFiles(b bool) C2EOption {
	return func(c *ChunkToExport) {
		c.includeFiles = b
	}
}

// WithSrcFileLoc sets the SrcFileLoc function.
func WithSrcFileLoc(fn func(*slack.Channel, *slack.File) string) C2EOption {
	return func(c *ChunkToExport) {
		if fn != nil {
			c.srcFileLoc = fn
		}
	}
}

// WithTrgFileLoc sets the TrgFileLoc function.
func WithTrgFileLoc(fn func(*slack.Channel, *slack.File) string) C2EOption {
	return func(c *ChunkToExport) {
		if fn != nil {
			c.trgFileLoc = fn
		}
	}
}

// WithLogger sets the logger.
func WithLogger(lg *slog.Logger) C2EOption {
	return func(c *ChunkToExport) {
		if lg != nil {
			c.lg = lg
		}
	}
}

func NewChunkToExport(src *chunk.Directory, trg fsadapter.FS, opt ...C2EOption) *ChunkToExport {
	c := &ChunkToExport{
		src:          src,
		trg:          trg,
		includeFiles: false,
		srcFileLoc:   fileproc.MattermostFilepath,
		trgFileLoc:   fileproc.MattermostFilepath,
		lg:           slog.Default(),
		request:      make(chan copyrequest, 1),
		result:       make(chan copyresult, 1),
		workers:      defWorkers,
	}
	for _, o := range opt {
		o(c)
	}
	return c
}

// Validate validates the input parameters.
func (c *ChunkToExport) Validate() error {
	if c.src == nil || c.trg == nil {
		return errors.New("convert: source and target must be set")
	}
	if c.includeFiles {
		const format = "convert: internal error: %s: %w"
		if c.srcFileLoc == nil {
			return fmt.Errorf(format, "source", ErrNoLocFunction)
		}
		if c.trgFileLoc == nil {
			return fmt.Errorf(format, "target", ErrNoLocFunction)
		}
	}
	// users chunk is required
	if fi, err := c.src.Stat(chunk.FUsers); err != nil {
		return fmt.Errorf("users chunk: %w", err)
	} else if fi.Size() == 0 {
		return fmt.Errorf("users chunk: %w", ErrEmptyChunk)
	}
	// we are not checking for channels chunk because channel information will
	// be stored in each of the chunk files, and we can collect it from there.

	return nil
}

// Convert converts the chunk directory contents to the export format.
// It validates the input parameters.
//
// # Restrictions
//
// Currently, one chunk file per channel is supported.  If there are multiple
// chunk files per channel, the behaviour is undefined, but I expect it to
// overwrite the previous files.
func (c *ChunkToExport) Convert(ctx context.Context) error {
	ctx, task := trace.NewTask(ctx, "convert.ChunkToExport")
	defer task.End()

	if err := c.Validate(); err != nil {
		return err
	}
	channels, err := c.src.Channels()
	if err != nil {
		return err
	}
	users, err := c.src.Users()
	if err != nil {
		return err
	}
	tfopts := []transform.ExpCvtOption{
		transform.ExpWithUsers(users),
	}
	if c.includeFiles {
		tfopts = append(tfopts, transform.ExpWithMsgUpdateFunc(func(ch *slack.Channel, m *slack.Message) error {
			// copy in a separate goroutine to avoid blocking the transform in
			// case of a synchronous fsadapter (e.g. zip file adapter can
			// write only one file at a time).
			c.request <- copyrequest{
				channel: ch,
				message: m,
			}
			return nil
		}))
		go c.copyworker(c.result, c.request)
	}

	// 1. generator
	chC := make(chan slack.Channel)
	go func() {
		defer close(chC)
		for _, ch := range channels {
			chC <- ch
		}
	}()

	errC := make(chan error, c.workers)
	{
		// 2. workers
		// 2.1 converter
		conv := transform.NewExpConverter(c.src, c.trg, tfopts...)
		var wg sync.WaitGroup
		for i := 0; i < c.workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for ch := range chC {
					lg := c.lg.With("channel", ch.ID)
					lg.Debug("processing channel")
					if err := conv.Convert(ctx, chunk.ToFileID(ch.ID, "", false)); err != nil {
						errC <- fmt.Errorf("converter: failed to process %q: %w", ch.ID, err)
						return
					}
				}
			}()
		}
		// 2.2 index writer
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.lg.DebugContext(ctx, "writing index", "name", c.src.Name())
			if err := conv.WriteIndex(); err != nil {
				errC <- err
			}
		}()
		// 2.3. workers sentinel
		go func() {
			wg.Wait()
			close(c.request)
		}()
	}

	// 3. result processor
LOOP:
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errC: // get rid of this shit.
			if err != nil {
				return err
			}
		case res, more := <-c.result:
			if !more {
				break LOOP
			}
			if res.err != nil {
				return fmt.Errorf("error processing message with ts=%s: %w", res.fr.message.Timestamp, res.err)
			}
		}
	}

	return nil
}

type copyerror struct {
	FileID string
	Err    error
}

func (e *copyerror) Error() string {
	return fmt.Sprintf("copy error: file ID=%s: %v", e.FileID, e.Err)
}

func (e *copyerror) Unwrap() error {
	return e.Err
}

// fileCopy iterates through the files in the message and copies them to the
// target directory.  Source file location is determined by calling the
// srcFileLoc function, joined with the chunk directory name.  target file
// location — by calling trgFileLoc function, and is relative to the target
// fsadapter root.
func (c *ChunkToExport) fileCopy(ch *slack.Channel, msg *slack.Message) error {
	if !c.includeFiles {
		return nil
	}
	if msg == nil {
		return errors.New("convert: internal error: callback: nil message")
	}
	if len(msg.Files) == 0 {
		// no files to process
		return nil
	}
	for _, f := range msg.Files {
		if err := fileproc.IsValidWithReason(&f); err != nil {
			c.lg.Warn("skipping", "file", f.ID, "error", err)
			continue
		}

		srcpath := filepath.Join(c.src.Name(), c.srcFileLoc(ch, &f))
		trgpath := c.trgFileLoc(ch, &f)

		if _, err := os.Stat(srcpath); err != nil {
			return &copyerror{f.ID, err}
		}
		if _, err := os.Stat(srcpath); err != nil {
			return &copyerror{f.ID, err}
		}
		c.lg.Debug("copying", "srcpath", srcpath, "trgpath", trgpath)
		if err := copy2trg(c.trg, trgpath, srcpath); err != nil {
			return &copyerror{f.ID, err}
		}
	}
	return nil
}

// copy2trg copies the file from the source path to the target path.  Source
// path is absolute, target path is relative to the target FS adapter root.
func copy2trg(trgfs fsadapter.FS, trgpath, srcpath string) error {
	in, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := trgfs.Create(trgpath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

type copyrequest struct {
	channel *slack.Channel
	message *slack.Message
}

type copyresult struct {
	fr  copyrequest
	err error
}

func (cr copyresult) Error() string {
	return fmt.Sprintf("copy error: channel [%s]: %s", cr.fr.channel.ID, cr.err)
}

func (cr copyresult) Unwrap() error {
	return cr.err
}

func (c *ChunkToExport) copyworker(res chan<- copyresult, req <-chan copyrequest) {
	defer close(res)
	for fr := range req {
		res <- copyresult{
			fr:  fr,
			err: c.fileCopy(fr.channel, fr.message),
		}
	}
}
