package executors

import "time"

const defaultChunkSize = 1024 * 1024

type (
	ChunkOption func(options *chunkOptions)

	ChunkExecutor struct {
		executor  *PeriodicalExecutor
		container *chunkContainer
	}

	chunkOptions struct {
		chunkSize     int
		flushInterval time.Duration
	}
)

func NewChunkExecutor(execute Execute, opts ...ChunkOption) *ChunkExecutor {
	_ = "STUB: not implemented"
	return nil
}

func (ce *ChunkExecutor) Add(task any, size int) error { _ = "STUB: not implemented"; return nil }

func (ce *ChunkExecutor) Flush() { _ = "STUB: not implemented"; return }

func (ce *ChunkExecutor) Wait() { _ = "STUB: not implemented"; return }

func WithChunkBytes(size int) ChunkOption { _ = "STUB: not implemented"; return *new(ChunkOption) }

func WithFlushInterval(duration time.Duration) ChunkOption {
	_ = "STUB: not implemented"
	return *new(ChunkOption)
}

func newChunkOptions() chunkOptions { _ = "STUB: not implemented"; return *new(chunkOptions) }

type chunkContainer struct {
	tasks        []any
	execute      Execute
	size         int
	maxChunkSize int
}

func (bc *chunkContainer) AddTask(task any) bool { _ = "STUB: not implemented"; return false }

func (bc *chunkContainer) Execute(tasks any) { _ = "STUB: not implemented"; return }

func (bc *chunkContainer) RemoveAll() any { _ = "STUB: not implemented"; return *new(any) }

type chunk struct {
	val  any
	size int
}
