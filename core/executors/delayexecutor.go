package executors

import (
	"sync"
	"time"
)

type DelayExecutor struct {
	fn        func()
	delay     time.Duration
	triggered bool
	lock      sync.Mutex
}

func NewDelayExecutor(fn func(), delay time.Duration) *DelayExecutor {
	_ = "STUB: not implemented"
	return nil
}

func (de *DelayExecutor) Trigger() { _ = "STUB: not implemented"; return }
