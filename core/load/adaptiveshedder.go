package load

import (
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/core/syncx"
)

const (
	defaultBuckets = 50
	defaultWindow  = time.Second * 5

	defaultCpuThreshold = 900
	defaultMinRt        = float64(time.Second / time.Millisecond)

	flyingBeta               = 0.9
	coolOffDuration          = time.Second
	cpuMax                   = 1000
	millisecondsPerSecond    = 1000
	overloadFactorLowerBound = 0.1
)

var (
	ErrServiceOverloaded = errors.New("service overloaded")

	enabled = syncx.ForAtomicBool(true)

	logEnabled = syncx.ForAtomicBool(true)

	systemOverloadChecker = func(cpuThreshold int64) bool {
		return stat.CpuUsage() >= cpuThreshold
	}
)

type (
	Promise interface {
		Pass()

		Fail()
	}

	Shedder interface {
		Allow() (Promise, error)
	}

	ShedderOption func(opts *shedderOptions)

	shedderOptions struct {
		window       time.Duration
		buckets      int
		cpuThreshold int64
	}

	adaptiveShedder struct {
		cpuThreshold    int64
		windowScale     float64
		flying          int64
		avgFlying       float64
		avgFlyingLock   syncx.SpinLock
		overloadTime    *syncx.AtomicDuration
		droppedRecently *syncx.AtomicBool
		passCounter     *collection.RollingWindow[int64, *collection.Bucket[int64]]
		rtCounter       *collection.RollingWindow[int64, *collection.Bucket[int64]]
	}
)

func Disable() { _ = "STUB: not implemented"; return }

func DisableLog() { _ = "STUB: not implemented"; return }

func NewAdaptiveShedder(opts ...ShedderOption) Shedder {
	_ = "STUB: not implemented"
	return *new(Shedder)
}

func (as *adaptiveShedder) Allow() (Promise, error) {
	_ = "STUB: not implemented"
	return *new(Promise), nil
}

func (as *adaptiveShedder) addFlying(delta int64) { _ = "STUB: not implemented"; return }

func (as *adaptiveShedder) highThru() bool { _ = "STUB: not implemented"; return false }

func (as *adaptiveShedder) maxFlight() float64 { _ = "STUB: not implemented"; return 0 }

func (as *adaptiveShedder) maxPass() int64 { _ = "STUB: not implemented"; return 0 }

func (as *adaptiveShedder) minRt() float64 { _ = "STUB: not implemented"; return 0 }

func (as *adaptiveShedder) overloadFactor() float64 { _ = "STUB: not implemented"; return 0 }

func (as *adaptiveShedder) shouldDrop() bool { _ = "STUB: not implemented"; return false }

func (as *adaptiveShedder) stillHot() bool { _ = "STUB: not implemented"; return false }

func (as *adaptiveShedder) systemOverloaded() bool { _ = "STUB: not implemented"; return false }

func WithBuckets(buckets int) ShedderOption { _ = "STUB: not implemented"; return *new(ShedderOption) }

func WithCpuThreshold(threshold int64) ShedderOption {
	_ = "STUB: not implemented"
	return *new(ShedderOption)
}

func WithWindow(window time.Duration) ShedderOption {
	_ = "STUB: not implemented"
	return *new(ShedderOption)
}

type promise struct {
	start   time.Duration
	shedder *adaptiveShedder
}

func (p *promise) Fail() { _ = "STUB: not implemented"; return }

func (p *promise) Pass() { _ = "STUB: not implemented"; return }
