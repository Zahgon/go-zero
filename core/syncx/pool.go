package syncx

import (
	"sync"
	"time"
)

type (
	PoolOption func(*Pool)

	node struct {
		item     any
		next     *node
		lastUsed time.Duration
	}

	Pool struct {
		limit   int
		created int
		maxAge  time.Duration
		lock    sync.Locker
		cond    *sync.Cond
		head    *node
		create  func() any
		destroy func(any)
	}
)

func NewPool(n int, create func() any, destroy func(any), opts ...PoolOption) *Pool {
	_ = "STUB: not implemented"
	return nil
}

func (p *Pool) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (p *Pool) Put(x any) { _ = "STUB: not implemented"; return }

func WithMaxAge(duration time.Duration) PoolOption {
	_ = "STUB: not implemented"
	return *new(PoolOption)
}
