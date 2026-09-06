package executors

import "time"

const defaultBulkTasks = 1000

type (
	BulkOption func(options *bulkOptions)

	BulkExecutor struct {
		executor  *PeriodicalExecutor
		container *bulkContainer
	}

	bulkOptions struct {
		cachedTasks   int
		flushInterval time.Duration
	}
)

func NewBulkExecutor(execute Execute, opts ...BulkOption) *BulkExecutor {
	_ = "STUB: not implemented"
	return nil
}

func (be *BulkExecutor) Add(task any) error { _ = "STUB: not implemented"; return nil }

func (be *BulkExecutor) Flush() { _ = "STUB: not implemented"; return }

func (be *BulkExecutor) Wait() { _ = "STUB: not implemented"; return }

func WithBulkTasks(tasks int) BulkOption { _ = "STUB: not implemented"; return *new(BulkOption) }

func WithBulkInterval(duration time.Duration) BulkOption {
	_ = "STUB: not implemented"
	return *new(BulkOption)
}

func newBulkOptions() bulkOptions { _ = "STUB: not implemented"; return *new(bulkOptions) }

type bulkContainer struct {
	tasks    []any
	execute  Execute
	maxTasks int
}

func (bc *bulkContainer) AddTask(task any) bool { _ = "STUB: not implemented"; return false }

func (bc *bulkContainer) Execute(tasks any) { _ = "STUB: not implemented"; return }

func (bc *bulkContainer) RemoveAll() any { _ = "STUB: not implemented"; return *new(any) }
