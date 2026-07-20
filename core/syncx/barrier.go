package syncx

import "sync"

type Barrier struct {
	lock sync.Mutex
}

func (b *Barrier) Guard(fn func()) { _ = "STUB: not implemented"; return }

func Guard(lock sync.Locker, fn func()) { _ = "STUB: not implemented"; return }
