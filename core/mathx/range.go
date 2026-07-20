package mathx

type Numerical interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func AtLeast[T Numerical](x, lower T) T { _ = "STUB: not implemented"; return *new(T) }

func AtMost[T Numerical](x, upper T) T { _ = "STUB: not implemented"; return *new(T) }

func Between[T Numerical](x, lower, upper T) T { _ = "STUB: not implemented"; return *new(T) }
