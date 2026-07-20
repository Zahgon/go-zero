package syncx

import (
	"time"

	"github.com/zeromicro/go-zero/core/lang"
)

type Cond struct {
	signal chan lang.PlaceholderType
}

func NewCond() *Cond { _ = "STUB: not implemented"; return nil }

func (cond *Cond) WaitWithTimeout(timeout time.Duration) (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}

func (cond *Cond) Wait() { _ = "STUB: not implemented"; return }

func (cond *Cond) Signal() { _ = "STUB: not implemented"; return }
