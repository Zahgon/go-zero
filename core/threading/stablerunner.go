package threading

import (
	"errors"
	"runtime"
	"sync"
)

const factor = 10

var (
	ErrRunnerClosed = errors.New("runner closed")

	bufSize = runtime.NumCPU() * factor
)

type StableRunner[I, O any] struct {
	handle        func(I) O
	consumedIndex uint64
	writtenIndex  uint64
	ring          []*struct {
		value chan O
		lock  sync.Mutex
	}
	runner *TaskRunner
	done   chan struct{}
}

func NewStableRunner[I, O any](fn func(I) O) *StableRunner[I, O] {
	_ = "STUB: not implemented"
	return nil
}

func (r *StableRunner[I, O]) Get() (O, error) { _ = "STUB: not implemented"; return *new(O), nil }

func (r *StableRunner[I, O]) Push(v I) error { _ = "STUB: not implemented"; return nil }

func (r *StableRunner[I, O]) Wait() { _ = "STUB: not implemented"; return }
