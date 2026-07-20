package syncx

import (
	"errors"
	"sync"
)

var ErrUseOfCleaned = errors.New("using a cleaned resource")

type RefResource struct {
	lock    sync.Mutex
	ref     int32
	cleaned bool
	clean   func()
}

func NewRefResource(clean func()) *RefResource { _ = "STUB: not implemented"; return nil }

func (r *RefResource) Use() error { _ = "STUB: not implemented"; return nil }

func (r *RefResource) Clean() { _ = "STUB: not implemented"; return }
