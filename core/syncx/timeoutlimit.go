package syncx

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("borrow timeout")

type TimeoutLimit struct {
	limit Limit
	cond  *Cond
}

func NewTimeoutLimit(n int) TimeoutLimit { _ = "STUB: not implemented"; return *new(TimeoutLimit) }

func (l TimeoutLimit) Borrow(timeout time.Duration) error { _ = "STUB: not implemented"; return nil }

func (l TimeoutLimit) Return() error { _ = "STUB: not implemented"; return nil }

func (l TimeoutLimit) TryBorrow() bool { _ = "STUB: not implemented"; return false }
