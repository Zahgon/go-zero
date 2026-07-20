package logx

import (
	"context"
	"sync"
	"sync/atomic"
)

var (
	globalFields     atomic.Value
	globalFieldsLock sync.Mutex
)

type fieldsKey struct{}

func AddGlobalFields(fields ...LogField) { _ = "STUB: not implemented"; return }

func ContextWithFields(ctx context.Context, fields ...LogField) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func WithFields(ctx context.Context, fields ...LogField) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
