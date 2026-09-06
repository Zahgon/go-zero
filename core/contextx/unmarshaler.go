package contextx

import (
	"context"

	"github.com/zeromicro/go-zero/core/mapping"
)

const contextTagKey = "ctx"

var unmarshaler = mapping.NewUnmarshaler(contextTagKey)

type contextValuer struct {
	context.Context
}

func (cv contextValuer) Value(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func For(ctx context.Context, v any) error { _ = "STUB: not implemented"; return nil }
