package executors

import (
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/lang"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/core/timex"
)

const idleRound = 10

type (
	TaskContainer interface {
		AddTask(task any) bool

		Execute(tasks any)

		RemoveAll() any
	}

	PeriodicalExecutor struct {
		commander chan any
		interval  time.Duration
		container TaskContainer
		waitGroup sync.WaitGroup

		wgBarrier   syncx.Barrier
		confirmChan chan lang.PlaceholderType
		inflight    int32
		guarded     bool
		newTicker   func(duration time.Duration) timex.Ticker
		lock        sync.Mutex
	}
)

func NewPeriodicalExecutor(interval time.Duration, container TaskContainer) *PeriodicalExecutor {
	_ = "STUB: not implemented"
	return nil
}

func (pe *PeriodicalExecutor) Add(task any) { _ = "STUB: not implemented"; return }

func (pe *PeriodicalExecutor) Flush() bool { _ = "STUB: not implemented"; return false }

func (pe *PeriodicalExecutor) Sync(fn func()) { _ = "STUB: not implemented"; return }

func (pe *PeriodicalExecutor) Wait() { _ = "STUB: not implemented"; return }

func (pe *PeriodicalExecutor) addAndCheck(task any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (pe *PeriodicalExecutor) backgroundFlush() { _ = "STUB: not implemented"; return }

func (pe *PeriodicalExecutor) doneExecution() { _ = "STUB: not implemented"; return }

func (pe *PeriodicalExecutor) enterExecution() { _ = "STUB: not implemented"; return }

func (pe *PeriodicalExecutor) executeTasks(tasks any) bool { _ = "STUB: not implemented"; return false }

func (pe *PeriodicalExecutor) hasTasks(tasks any) bool { _ = "STUB: not implemented"; return false }

func (pe *PeriodicalExecutor) shallQuit(last time.Duration) (stop bool) {
	_ = "STUB: not implemented"
	return false
}
