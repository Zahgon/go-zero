package collection

import (
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/mathx"
)

type (
	BucketInterface[T Numerical] interface {
		Add(v T)
		Reset()
	}

	Numerical = mathx.Numerical

	RollingWindowOption[T Numerical, B BucketInterface[T]] func(rollingWindow *RollingWindow[T, B])

	RollingWindow[T Numerical, B BucketInterface[T]] struct {
		lock          sync.RWMutex
		size          int
		win           *window[T, B]
		interval      time.Duration
		offset        int
		ignoreCurrent bool
		lastTime      time.Duration
	}
)

func NewRollingWindow[T Numerical, B BucketInterface[T]](newBucket func() B, size int,
	interval time.Duration, opts ...RollingWindowOption[T, B]) *RollingWindow[T, B] {
	_ = "STUB: not implemented"
	return nil
}

func (rw *RollingWindow[T, B]) Add(v T) { _ = "STUB: not implemented"; return }

func (rw *RollingWindow[T, B]) Reduce(fn func(b B)) { _ = "STUB: not implemented"; return }

func (rw *RollingWindow[T, B]) span() int { _ = "STUB: not implemented"; return 0 }

func (rw *RollingWindow[T, B]) updateOffset() { _ = "STUB: not implemented"; return }

type Bucket[T Numerical] struct {
	Sum   T
	Count int64
}

func (b *Bucket[T]) Add(v T) { _ = "STUB: not implemented"; return }

func (b *Bucket[T]) Reset() { _ = "STUB: not implemented"; return }

type window[T Numerical, B BucketInterface[T]] struct {
	buckets []B
	size    int
}

func newWindow[T Numerical, B BucketInterface[T]](newBucket func() B, size int) *window[T, B] {
	_ = "STUB: not implemented"
	return nil
}

func (w *window[T, B]) add(offset int, v T) { _ = "STUB: not implemented"; return }

func (w *window[T, B]) reduce(start, count int, fn func(b B)) { _ = "STUB: not implemented"; return }

func (w *window[T, B]) resetBucket(offset int) { _ = "STUB: not implemented"; return }

func IgnoreCurrentBucket[T Numerical, B BucketInterface[T]]() RollingWindowOption[T, B] {
	_ = "STUB: not implemented"
	return nil
}
