package mr

import (
	"context"
	"errors"
)

const (
	defaultWorkers = 16
	minWorkers     = 1
)

var (
	ErrCancelWithNil = errors.New("mapreduce cancelled with nil")

	ErrReduceNoOutput = errors.New("reduce not writing value")
)

type (
	ForEachFunc[T any] func(item T)

	GenerateFunc[T any] func(source chan<- T)

	MapFunc[T, U any] func(item T, writer Writer[U])

	MapperFunc[T, U any] func(item T, writer Writer[U], cancel func(error))

	ReducerFunc[U, V any] func(pipe <-chan U, writer Writer[V], cancel func(error))

	VoidReducerFunc[U any] func(pipe <-chan U, cancel func(error))

	Option func(opts *mapReduceOptions)

	mapperContext[T, U any] struct {
		ctx       context.Context
		mapper    MapFunc[T, U]
		source    <-chan T
		panicChan *onceChan
		collector chan<- U
		doneChan  <-chan struct{}
		workers   int
	}

	mapReduceOptions struct {
		ctx     context.Context
		workers int
	}

	Writer[T any] interface {
		Write(v T)
	}
)

func Finish(fns ...func() error) error { _ = "STUB: not implemented"; return nil }

func FinishVoid(fns ...func()) { _ = "STUB: not implemented"; return }

func ForEach[T any](generate GenerateFunc[T], mapper ForEachFunc[T], opts ...Option) {
	_ = "STUB: not implemented"
	return
}

func MapReduce[T, U, V any](generate GenerateFunc[T], mapper MapperFunc[T, U], reducer ReducerFunc[U, V],
	opts ...Option) (V, error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}

func MapReduceChan[T, U, V any](source <-chan T, mapper MapperFunc[T, U], reducer ReducerFunc[U, V],
	opts ...Option) (V, error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}

func MapReduceVoid[T, U any](generate GenerateFunc[T], mapper MapperFunc[T, U],
	reducer VoidReducerFunc[U], opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWorkers(workers int) Option { _ = "STUB: not implemented"; return *new(Option) }

func buildOptions(opts ...Option) *mapReduceOptions { _ = "STUB: not implemented"; return nil }

func buildPanicInfo(r any, stack []byte) string { _ = "STUB: not implemented"; return "" }

func buildSource[T any](generate GenerateFunc[T], panicChan *onceChan) chan T {
	_ = "STUB: not implemented"
	return nil
}

func drain[T any](channel <-chan T) { _ = "STUB: not implemented"; return }

func executeMappers[T, U any](mCtx mapperContext[T, U]) { _ = "STUB: not implemented"; return }

func mapReduceWithPanicChan[T, U, V any](source <-chan T, panicChan *onceChan, mapper MapperFunc[T, U],
	reducer ReducerFunc[U, V], opts ...Option) (val V, err error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}

func newOptions() *mapReduceOptions { _ = "STUB: not implemented"; return nil }

func once(fn func(error)) func(error) { _ = "STUB: not implemented"; return nil }

type guardedWriter[T any] struct {
	ctx     context.Context
	channel chan<- T
	done    <-chan struct{}
}

func newGuardedWriter[T any](ctx context.Context, channel chan<- T, done <-chan struct{}) guardedWriter[T] {
	_ = "STUB: not implemented"
	return nil
}

func (gw guardedWriter[T]) Write(v T) { _ = "STUB: not implemented"; return }

type onceChan struct {
	channel chan any
	wrote   int32
}

func (oc *onceChan) write(val any) { _ = "STUB: not implemented"; return }
