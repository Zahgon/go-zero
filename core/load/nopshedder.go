package load

type nopShedder struct{}

func newNopShedder() Shedder { _ = "STUB: not implemented"; return *new(Shedder) }

func (s nopShedder) Allow() (Promise, error) { _ = "STUB: not implemented"; return *new(Promise), nil }

type nopPromise struct{}

func (p nopPromise) Pass() { _ = "STUB: not implemented"; return }

func (p nopPromise) Fail() { _ = "STUB: not implemented"; return }
