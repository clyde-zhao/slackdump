package expproc

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rusq/fsadapter"
	"github.com/rusq/slackdump/v2/internal/chunk"
	"github.com/rusq/slackdump/v2/internal/osext"
	"github.com/slack-go/slack"
)

var errNotADir = errors.New("not a directory")

type Transform struct {
	srcdir string       // source directory with chunks.
	fsa    fsadapter.FS // target file system adapter.
	ids    chan string  // channel used to pass channel IDs to the worker.
	err    chan error   // error channel used to propagate errors to the main thread.
}

// NewTransform creates a new Transform instance.
func NewTransform(fsa fsadapter.FS, chunkdir string) (*Transform, error) {
	t := &Transform{
		srcdir: chunkdir,
		fsa:    fsa,
		ids:    make(chan string),
		err:    make(chan error, 1),
	}
	go t.worker()
	return t, nil
}

// OnFinish is the function that should be passed to the Channel processor.
func (t *Transform) OnFinish(channelID string) error {
	select {
	case err := <-t.err:
		return err
	default:
		t.ids <- channelID
	}
	return nil
}

func (t *Transform) worker() {
	for id := range t.ids {
		if err := mmtransform(t.fsa, t.srcdir, id); err != nil {
			t.err <- err
			continue
		}
	}
}

// Close closes the Transform processor.  It must be called once it is
// guaranteed that OnFinish will not be called anymore, otherwise the
// call to OnFinish will panic.
func (t *Transform) Close() error {
	close(t.ids)
	return nil
}

// mmtransform is the mattermost transformer.  It transforms the chunk file
// for the channel with ID into a slack export format, and attachments are
// placed into the __uploads__ directory which is understood by the mattermost
// import tool.  It expects the chunk file to be in the srcdir/id.json.gz
// file, and the attachments to be in the srcdir/id directory.
func mmtransform(fsa fsadapter.FS, srcdir string, id string) error {
	// load the chunk file
	f, err := openChunks(filepath.Join(srcdir, id+ext))
	if err != nil {
		return err
	}
	defer f.Close()
	// locate attachments
	// var hasAttachments bool
	// if hasAttachments, err = dirExists(filepath.Join(srcdir, id)); err != nil {
	// 	return err
	// }
	// transform the chunk file
	pl, err := chunk.NewPlayer(f)
	if err != nil {
		return err
	}

	ci, err := pl.ChannelInfo(id)
	if err != nil {
		return err
	}

	if err := writeMessages(fsa, ci, pl); err != nil {
		return err
	}

	return nil
}

func channelName(ch *slack.Channel) string {
	if ch.IsIM {
		return ch.ID
	}
	return ch.Name
}

func writeMessages(fsa fsadapter.FS, ci *slack.Channel, pl *chunk.Player) error {
	dir := channelName(ci)
	var prevDt string
	var wc io.WriteCloser
	var enc *json.Encoder
	if err := pl.Sorted(func(ts time.Time, m *slack.Message) error {
		date := ts.Format("2006-01-02")
		if date != prevDt || prevDt == "" {
			if wc != nil {
				if err := writeJSONFooter(wc); err != nil {
					return err
				}
				if err := wc.Close(); err != nil {
					return err
				}
			}
			var err error
			wc, err = fsa.Create(filepath.Join(dir, date+".json"))
			if err != nil {
				return err
			}
			if err := writeJSONHeader(wc); err != nil {
				return err
			}
			prevDt = date
			enc = json.NewEncoder(wc)
			enc.SetIndent("", "  ")
		} else {
			wc.Write([]byte(",\n"))
		}
		return enc.Encode(m)
	}); err != nil {
		return err
	}
	// write the last footer
	if wc != nil {
		if err := writeJSONFooter(wc); err != nil {
			return err
		}
		if err := wc.Close(); err != nil {
			return err
		}
	}
	return nil
}

func writeJSONHeader(w io.Writer) error {
	_, err := io.WriteString(w, "[\n")
	return err
}

func writeJSONFooter(w io.Writer) error {
	_, err := io.WriteString(w, "\n]\n")
	return err
}

func openChunks(filename string) (io.ReadSeekCloser, error) {
	if fi, err := os.Stat(filename); err != nil {
		return nil, err
	} else if fi.IsDir() {
		return nil, errors.New("chunk file is a directory")
	} else if fi.Size() == 0 {
		return nil, errors.New("chunk file is empty")
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tf, err := osext.UnGZIP(f)
	if err != nil {
		return nil, err
	}

	return osext.RemoveOnClose(tf, tf.Name()), nil
}

func dirExists(dir string) (bool, error) {
	fi, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if !fi.IsDir() {
		return false, errNotADir
	}
	return true, nil
}
