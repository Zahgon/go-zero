package syncx

import "sync"

type ManagedResource struct {
	resource any
	lock     sync.RWMutex
	generate func() any
	equals   func(a, b any) bool
}

func NewManagedResource(generate func() any, equals func(a, b any) bool) *ManagedResource {
	_ = "STUB: not implemented"
	return nil
}

func (mr *ManagedResource) MarkBroken(resource any) { _ = "STUB: not implemented"; return }

func (mr *ManagedResource) Take() any { _ = "STUB: not implemented"; return *new(any) }
