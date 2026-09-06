package syncx

type AtomicFloat64 uint64

func NewAtomicFloat64() *AtomicFloat64 { _ = "STUB: not implemented"; return nil }

func ForAtomicFloat64(val float64) *AtomicFloat64 { _ = "STUB: not implemented"; return nil }

func (f *AtomicFloat64) Add(val float64) float64 { _ = "STUB: not implemented"; return 0 }

func (f *AtomicFloat64) CompareAndSwap(old, val float64) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *AtomicFloat64) Load() float64 { _ = "STUB: not implemented"; return 0 }

func (f *AtomicFloat64) Set(val float64) { _ = "STUB: not implemented"; return }
