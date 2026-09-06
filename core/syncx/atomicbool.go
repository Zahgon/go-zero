package syncx

type AtomicBool uint32

func NewAtomicBool() *AtomicBool { _ = "STUB: not implemented"; return nil }

func ForAtomicBool(val bool) *AtomicBool { _ = "STUB: not implemented"; return nil }

func (b *AtomicBool) CompareAndSwap(old, val bool) bool { _ = "STUB: not implemented"; return false }

func (b *AtomicBool) Set(v bool) { _ = "STUB: not implemented"; return }

func (b *AtomicBool) True() bool { _ = "STUB: not implemented"; return false }
