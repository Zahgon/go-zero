package syncx

import (
	"sync"
	"time"
)

const defaultRefreshInterval = time.Second

type (
	ImmutableResourceOption func(resource *ImmutableResource)

	ImmutableResource struct {
		fetch           func() (any, error)
		resource        any
		err             error
		lock            sync.RWMutex
		refreshInterval time.Duration
		lastTime        *AtomicDuration
	}
)

func NewImmutableResource(fn func() (any, error), opts ...ImmutableResourceOption) *ImmutableResource {
	_ = "STUB: not implemented"
	return nil
}

func (ir *ImmutableResource) Get() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (ir *ImmutableResource) shouldRefresh() bool { _ = "STUB: not implemented"; return false }

func WithRefreshIntervalOnFailure(interval time.Duration) ImmutableResourceOption {
	_ = "STUB: not implemented"
	return *new(ImmutableResourceOption)
}
