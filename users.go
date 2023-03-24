package slackdump

// In this file: user related code.

import (
	"context"
	"errors"
	"runtime/trace"
	"time"

	"github.com/slack-go/slack"

	"github.com/rusq/slackdump/v2/internal/network"
	"github.com/rusq/slackdump/v2/types"
)

const (
	cacheTimeout = 5 * time.Minute
)

// GetUsers retrieves all users either from cache or from the API.  If
// Session.usercache is not empty, it will return the cached users.
// Otherwise, it will try fetching them from the API and cache them.
func (s *Session) GetUsers(ctx context.Context) (types.Users, error) {
	ctx, task := trace.NewTask(ctx, "GetUsers")
	defer task.End()

	// try getting users from cache
	users, err := s.uc.get(cacheTimeout)
	if err == nil {
		return users, nil
	}

	// if not succeeded, fetch them from the API.
	users, err = s.fetchUsers(ctx)
	if err != nil {
		return nil, err
	}
	s.uc.set(users)
	return users, err

}

// fetchUsers fetches users from the API.
func (s *Session) fetchUsers(ctx context.Context) (types.Users, error) {
	var (
		users []slack.User
	)
	l := network.NewLimiter(
		network.Tier2, s.cfg.limits.Tier2.Burst, int(s.cfg.limits.Tier2.Boost),
	)
	if err := withRetry(ctx, l, s.cfg.limits.Tier2.Retries, func() error {
		var err error
		users, err = s.client.GetUsersContext(ctx)
		return err
	}); err != nil {
		trace.Logf(ctx, "error", "GetUsers error=%s", err)
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("couldn't fetch users")
	}
	return users, nil
}
