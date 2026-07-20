package fx

import (
	"context"
	"time"
)

var (
	ErrCanceled = context.Canceled

	ErrTimeout = context.DeadlineExceeded
)

type DoOption func() context.Context

func DoWithTimeout(fn func() error, timeout time.Duration, opts ...DoOption) error {
	_ = "STUB: not implemented"
	return nil
}

func WithContext(ctx context.Context) DoOption { _ = "STUB: not implemented"; return *new(DoOption) }
