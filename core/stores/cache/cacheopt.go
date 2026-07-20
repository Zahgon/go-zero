package cache

import "time"

const (
	defaultExpiry         = time.Hour * 24 * 7
	defaultNotFoundExpiry = time.Minute
)

type (
	Options struct {
		Expiry         time.Duration
		NotFoundExpiry time.Duration
	}

	Option func(o *Options)
)

func newOptions(opts ...Option) Options { _ = "STUB: not implemented"; return *new(Options) }

func WithExpiry(expiry time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithNotFoundExpiry(expiry time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
