package errorx

import "sync/atomic"

type AtomicError struct {
	err atomic.Value
}

func (ae *AtomicError) Set(err error) { _ = "STUB: not implemented"; return }

func (ae *AtomicError) Load() error { _ = "STUB: not implemented"; return nil }
