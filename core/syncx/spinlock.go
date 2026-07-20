package syncx

type SpinLock struct {
	lock uint32
}

func (sl *SpinLock) Lock() { _ = "STUB: not implemented"; return }

func (sl *SpinLock) TryLock() bool { _ = "STUB: not implemented"; return false }

func (sl *SpinLock) Unlock() { _ = "STUB: not implemented"; return }
