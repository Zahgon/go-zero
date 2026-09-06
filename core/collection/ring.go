package collection

import "sync"

type Ring struct {
	elements []any
	index    int
	lock     sync.RWMutex
}

func NewRing(n int) *Ring { _ = "STUB: not implemented"; return nil }

func (r *Ring) Add(v any) { _ = "STUB: not implemented"; return }

func (r *Ring) Take() []any { _ = "STUB: not implemented"; return nil }
