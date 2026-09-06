package fx

const (
	defaultWorkers = 16
	minWorkers     = 1
)

type (
	rxOptions struct {
		unlimitedWorkers bool
		workers          int
	}

	FilterFunc func(item any) bool

	ForAllFunc func(pipe <-chan any)

	ForEachFunc func(item any)

	GenerateFunc func(source chan<- any)

	KeyFunc func(item any) any

	LessFunc func(a, b any) bool

	MapFunc func(item any) any

	Option func(opts *rxOptions)

	ParallelFunc func(item any)

	ReduceFunc func(pipe <-chan any) (any, error)

	WalkFunc func(item any, pipe chan<- any)

	Stream struct {
		source <-chan any
	}
)

func Concat(s Stream, others ...Stream) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func From(generate GenerateFunc) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func Just(items ...any) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func Range(source <-chan any) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) AllMatch(predicate func(item any) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (s Stream) AnyMatch(predicate func(item any) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (s Stream) Buffer(n int) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) Concat(others ...Stream) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) Count() (count int) { _ = "STUB: not implemented"; return 0 }

func (s Stream) Distinct(fn KeyFunc) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) Done() { _ = "STUB: not implemented"; return }

func (s Stream) Filter(fn FilterFunc, opts ...Option) Stream {
	_ = "STUB: not implemented"
	return *new(Stream)
}

func (s Stream) First() any { _ = "STUB: not implemented"; return *new(any) }

func (s Stream) ForAll(fn ForAllFunc) { _ = "STUB: not implemented"; return }

func (s Stream) ForEach(fn ForEachFunc) { _ = "STUB: not implemented"; return }

func (s Stream) Group(fn KeyFunc) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) Head(n int64) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) Last() (item any) { _ = "STUB: not implemented"; return *new(any) }

func (s Stream) Map(fn MapFunc, opts ...Option) Stream {
	_ = "STUB: not implemented"
	return *new(Stream)
}

func (s Stream) Max(less LessFunc) any { _ = "STUB: not implemented"; return *new(any) }

func (s Stream) Merge() Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) Min(less LessFunc) any { _ = "STUB: not implemented"; return *new(any) }

func (s Stream) NoneMatch(predicate func(item any) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (s Stream) Parallel(fn ParallelFunc, opts ...Option) { _ = "STUB: not implemented"; return }

func (s Stream) Reduce(fn ReduceFunc) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s Stream) Reverse() Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) Skip(n int64) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) Sort(less LessFunc) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) Split(n int) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) Tail(n int64) Stream { _ = "STUB: not implemented"; return *new(Stream) }

func (s Stream) Walk(fn WalkFunc, opts ...Option) Stream {
	_ = "STUB: not implemented"
	return *new(Stream)
}

func (s Stream) walkLimited(fn WalkFunc, option *rxOptions) Stream {
	_ = "STUB: not implemented"
	return *new(Stream)
}

func (s Stream) walkUnlimited(fn WalkFunc, option *rxOptions) Stream {
	_ = "STUB: not implemented"
	return *new(Stream)
}

func UnlimitedWorkers() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWorkers(workers int) Option { _ = "STUB: not implemented"; return *new(Option) }

func buildOptions(opts ...Option) *rxOptions { _ = "STUB: not implemented"; return nil }

func drain(channel <-chan any) { _ = "STUB: not implemented"; return }

func newOptions() *rxOptions { _ = "STUB: not implemented"; return nil }
