package syncx

import "sync"

type (
	LockedCalls interface {
		Do(key string, fn func() (any, error)) (any, error)
	}

	lockedGroup struct {
		mu sync.Mutex
		m  map[string]*sync.WaitGroup
	}
)

func NewLockedCalls() LockedCalls { _ = "STUB: not implemented"; return *new(LockedCalls) }

func (lg *lockedGroup) Do(key string, fn func() (any, error)) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (lg *lockedGroup) makeCall(key string, fn func() (any, error)) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
