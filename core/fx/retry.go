package fx

import (
	"context"
	"time"
)

const defaultRetryTimes = 3

type (
	RetryOption func(*retryOptions)

	retryOptions struct {
		times        int
		interval     time.Duration
		timeout      time.Duration
		ignoreErrors []error
	}
)

func DoWithRetry(fn func() error, opts ...RetryOption) error { _ = "STUB: not implemented"; return nil }

func DoWithRetryCtx(ctx context.Context, fn func(ctx context.Context, retryCount int) error,
	opts ...RetryOption) error {
	_ = "STUB: not implemented"
	return nil
}

func retry(ctx context.Context, fn func(errChan chan error, retryCount int), opts ...RetryOption) error {
	_ = "STUB: not implemented"
	return nil
}

func WithIgnoreErrors(ignoreErrors []error) RetryOption {
	_ = "STUB: not implemented"
	return *new(RetryOption)
}

func WithInterval(interval time.Duration) RetryOption {
	_ = "STUB: not implemented"
	return *new(RetryOption)
}

func WithRetry(times int) RetryOption { _ = "STUB: not implemented"; return *new(RetryOption) }

func WithTimeout(timeout time.Duration) RetryOption {
	_ = "STUB: not implemented"
	return *new(RetryOption)
}

func newRetryOptions() *retryOptions { _ = "STUB: not implemented"; return nil }
