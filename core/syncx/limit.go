package syncx

import (
	"errors"

	"github.com/zeromicro/go-zero/core/lang"
)

var ErrLimitReturn = errors.New("discarding limited token, resource pool is full, someone returned multiple times")

type Limit struct {
	pool chan lang.PlaceholderType
}

func NewLimit(n int) Limit { _ = "STUB: not implemented"; return *new(Limit) }

func (l Limit) Borrow() { _ = "STUB: not implemented"; return }

func (l Limit) Return() error { _ = "STUB: not implemented"; return nil }

func (l Limit) TryBorrow() bool { _ = "STUB: not implemented"; return false }
