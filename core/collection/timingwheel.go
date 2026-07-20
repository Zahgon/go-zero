package collection

import (
	"container/list"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/lang"
	"github.com/zeromicro/go-zero/core/timex"
)

const drainWorkers = 8

var (
	ErrClosed   = errors.New("TimingWheel is closed already")
	ErrArgument = errors.New("incorrect task argument")
)

type (
	Execute func(key, value any)

	TimingWheel struct {
		interval      time.Duration
		ticker        timex.Ticker
		slots         []*list.List
		timers        *SafeMap
		tickedPos     int
		numSlots      int
		execute       Execute
		setChannel    chan timingEntry
		moveChannel   chan baseEntry
		removeChannel chan any
		drainChannel  chan func(key, value any)
		stopChannel   chan lang.PlaceholderType
	}

	timingEntry struct {
		baseEntry
		value   any
		circle  int
		diff    int
		removed bool
	}

	baseEntry struct {
		delay time.Duration
		key   any
	}

	positionEntry struct {
		pos  int
		item *timingEntry
	}

	timingTask struct {
		key   any
		value any
	}
)

func NewTimingWheel(interval time.Duration, numSlots int, execute Execute) (*TimingWheel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTimingWheelWithTicker(interval time.Duration, numSlots int, execute Execute,
	ticker timex.Ticker) (*TimingWheel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tw *TimingWheel) Drain(fn func(key, value any)) error { _ = "STUB: not implemented"; return nil }

func (tw *TimingWheel) MoveTimer(key any, delay time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (tw *TimingWheel) RemoveTimer(key any) error { _ = "STUB: not implemented"; return nil }

func (tw *TimingWheel) SetTimer(key, value any, delay time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (tw *TimingWheel) Stop() { _ = "STUB: not implemented"; return }

func (tw *TimingWheel) drainAll(fn func(key, value any)) { _ = "STUB: not implemented"; return }

func (tw *TimingWheel) getPositionAndCircle(d time.Duration) (pos, circle int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (tw *TimingWheel) initSlots() { _ = "STUB: not implemented"; return }

func (tw *TimingWheel) moveTask(task baseEntry) { _ = "STUB: not implemented"; return }

func (tw *TimingWheel) onTick() { _ = "STUB: not implemented"; return }

func (tw *TimingWheel) removeTask(key any) { _ = "STUB: not implemented"; return }

func (tw *TimingWheel) run() { _ = "STUB: not implemented"; return }

func (tw *TimingWheel) runTasks(tasks []timingTask) { _ = "STUB: not implemented"; return }

func (tw *TimingWheel) scanAndRunTasks(l *list.List) { _ = "STUB: not implemented"; return }

func (tw *TimingWheel) setTask(task *timingEntry) { _ = "STUB: not implemented"; return }

func (tw *TimingWheel) setTimerPosition(pos int, task *timingEntry) {
	_ = "STUB: not implemented"
	return
}
