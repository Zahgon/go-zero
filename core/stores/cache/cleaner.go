package cache

import (
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/core/threading"
)

const (
	timingWheelSlots = 300
	cleanWorkers     = 5
	taskKeyLen       = 8
)

var (
	timingWheel atomic.Value
	taskRunner  = threading.NewTaskRunner(cleanWorkers)
)

type delayTask struct {
	delay time.Duration
	task  func() error
	keys  []string
}

func init() {
	tw, err := collection.NewTimingWheel(time.Second, timingWheelSlots, clean)
	logx.Must(err)
	timingWheel.Store(tw)

	proc.AddShutdownListener(func() {
		if err := tw.Drain(clean); err != nil {
			logx.Errorf("failed to drain timing wheel: %v", err)
		}
	})
}

func AddCleanTask(task func() error, keys ...string) { _ = "STUB: not implemented"; return }

func clean(key, value any) { _ = "STUB: not implemented"; return }

func nextDelay(delay time.Duration) (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}
