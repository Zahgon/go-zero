package errorx

import (
	"sync"
)

type BatchError struct {
	errs []error
	lock sync.RWMutex
}

func (be *BatchError) Add(errs ...error) { _ = "STUB: not implemented"; return }

func (be *BatchError) Err() error { _ = "STUB: not implemented"; return nil }

func (be *BatchError) NotNil() bool { _ = "STUB: not implemented"; return false }
