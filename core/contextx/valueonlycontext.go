package contextx

import (
	"context"
	"time"
)

type valueOnlyContext struct {
	context.Context
}

func (valueOnlyContext) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (valueOnlyContext) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (valueOnlyContext) Err() error { _ = "STUB: not implemented"; return nil }

func ValueOnlyFrom(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
