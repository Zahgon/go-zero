package threading

import (
	"errors"
	"sync"

	"github.com/zeromicro/go-zero/core/lang"
)

var ErrTaskRunnerBusy = errors.New("task runner is busy")

type TaskRunner struct {
	limitChan chan lang.PlaceholderType
	waitGroup sync.WaitGroup
}

func NewTaskRunner(concurrency int) *TaskRunner { _ = "STUB: not implemented"; return nil }

func (rp *TaskRunner) Schedule(task func()) { _ = "STUB: not implemented"; return }

func (rp *TaskRunner) ScheduleImmediately(task func()) error { _ = "STUB: not implemented"; return nil }

func (rp *TaskRunner) Wait() { _ = "STUB: not implemented"; return }
