package collection

import (
	"sync"
)

const (
	copyThreshold = 1000
	maxDeletion   = 10000
)

type SafeMap struct {
	lock        sync.RWMutex
	deletionOld int
	deletionNew int
	dirtyOld    map[any]any
	dirtyNew    map[any]any
}

func NewSafeMap() *SafeMap { _ = "STUB: not implemented"; return nil }

func (m *SafeMap) Del(key any) { _ = "STUB: not implemented"; return }

func (m *SafeMap) Get(key any) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

func (m *SafeMap) Range(f func(key, val any) bool) { _ = "STUB: not implemented"; return }

func (m *SafeMap) Set(key, value any) { _ = "STUB: not implemented"; return }

func (m *SafeMap) Size() int { _ = "STUB: not implemented"; return 0 }
