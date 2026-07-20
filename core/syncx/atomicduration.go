package syncx

import (
	"time"
)

type AtomicDuration int64

func NewAtomicDuration() *AtomicDuration { _ = "STUB: not implemented"; return nil }

func ForAtomicDuration(val time.Duration) *AtomicDuration { _ = "STUB: not implemented"; return nil }

func (d *AtomicDuration) CompareAndSwap(old, val time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *AtomicDuration) Load() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (d *AtomicDuration) Set(val time.Duration) { _ = "STUB: not implemented"; return }
