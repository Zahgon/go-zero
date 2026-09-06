package redis

import (
	"context"
	"time"

	red "github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/attribute"
)

const spanName = "redis"

var (
	defaultDurationHook   = durationHook{}
	redisCmdsAttributeKey = attribute.Key("redis.cmds")
)

type durationHook struct {
}

func (h durationHook) DialHook(next red.DialHook) red.DialHook {
	_ = "STUB: not implemented"
	return *new(red.DialHook)
}

func (h durationHook) ProcessHook(next red.ProcessHook) red.ProcessHook {
	_ = "STUB: not implemented"
	return *new(red.ProcessHook)
}

func (h durationHook) ProcessPipelineHook(next red.ProcessPipelineHook) red.ProcessPipelineHook {
	_ = "STUB: not implemented"
	return *new(red.ProcessPipelineHook)
}

func (h durationHook) startSpan(ctx context.Context, cmds ...red.Cmder) (context.Context, func(err error)) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func formatError(err error) string { _ = "STUB: not implemented"; return "" }

func logDuration(ctx context.Context, cmds []red.Cmder, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}
