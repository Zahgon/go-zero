package breaker

import (
	"time"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/mathx"
	"github.com/zeromicro/go-zero/core/syncx"
)

const (
	window            = time.Second * 10
	buckets           = 40
	forcePassDuration = time.Second
	k                 = 1.5
	minK              = 1.1
	protection        = 5
)

type (
	googleBreaker struct {
		k        float64
		stat     *collection.RollingWindow[int64, *bucket]
		proba    *mathx.Proba
		lastPass *syncx.AtomicDuration
	}

	windowResult struct {
		accepts        int64
		total          int64
		failingBuckets int64
		workingBuckets int64
	}
)

func newGoogleBreaker() *googleBreaker { _ = "STUB: not implemented"; return nil }

func (b *googleBreaker) accept() error { _ = "STUB: not implemented"; return nil }

func (b *googleBreaker) allow() (internalPromise, error) {
	_ = "STUB: not implemented"
	return *new(internalPromise), nil
}

func (b *googleBreaker) doReq(req func() error, fallback Fallback, acceptable Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *googleBreaker) markDrop() { _ = "STUB: not implemented"; return }

func (b *googleBreaker) markFailure() { _ = "STUB: not implemented"; return }

func (b *googleBreaker) markSuccess() { _ = "STUB: not implemented"; return }

func (b *googleBreaker) history() windowResult {
	_ = "STUB: not implemented"
	return *new(windowResult)
}

type googlePromise struct {
	b *googleBreaker
}

func (p googlePromise) Accept() { _ = "STUB: not implemented"; return }

func (p googlePromise) Reject() { _ = "STUB: not implemented"; return }
