package breaker

const (
	success = iota
	fail
	drop
)

type bucket struct {
	Sum     int64
	Success int64
	Failure int64
	Drop    int64
}

func (b *bucket) Add(v int64) { _ = "STUB: not implemented"; return }

func (b *bucket) Reset() { _ = "STUB: not implemented"; return }

func (b *bucket) drop() { _ = "STUB: not implemented"; return }

func (b *bucket) fail() { _ = "STUB: not implemented"; return }

func (b *bucket) succeed() { _ = "STUB: not implemented"; return }
