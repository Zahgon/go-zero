package mathx

import (
	"math/rand"
	"sync"
	"time"
)

type Unstable struct {
	deviation float64
	r         *rand.Rand
	lock      *sync.Mutex
}

func NewUnstable(deviation float64) Unstable { _ = "STUB: not implemented"; return *new(Unstable) }

func (u Unstable) AroundDuration(base time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (u Unstable) AroundInt(base int64) int64 { _ = "STUB: not implemented"; return 0 }
