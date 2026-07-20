package logc

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type (
	LogConf  = logx.LogConf
	LogField = logx.LogField
)

func AddGlobalFields(fields ...LogField) { _ = "STUB: not implemented"; return }

func Alert(_ context.Context, v string) { _ = "STUB: not implemented"; return }

func Close() error { _ = "STUB: not implemented"; return nil }

func Debug(ctx context.Context, v ...interface{}) { _ = "STUB: not implemented"; return }

func Debugf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func Debugfn(ctx context.Context, fn func() any) { _ = "STUB: not implemented"; return }

func Debugv(ctx context.Context, v interface{}) { _ = "STUB: not implemented"; return }

func Debugw(ctx context.Context, msg string, fields ...LogField) { _ = "STUB: not implemented"; return }

func Error(ctx context.Context, v ...any) { _ = "STUB: not implemented"; return }

func Errorf(ctx context.Context, format string, v ...any) { _ = "STUB: not implemented"; return }

func Errorfn(ctx context.Context, fn func() any) { _ = "STUB: not implemented"; return }

func Errorv(ctx context.Context, v any) { _ = "STUB: not implemented"; return }

func Errorw(ctx context.Context, msg string, fields ...LogField) { _ = "STUB: not implemented"; return }

func Field(key string, value any) LogField { _ = "STUB: not implemented"; return *new(LogField) }

func Info(ctx context.Context, v ...any) { _ = "STUB: not implemented"; return }

func Infof(ctx context.Context, format string, v ...any) { _ = "STUB: not implemented"; return }

func Infofn(ctx context.Context, fn func() any) { _ = "STUB: not implemented"; return }

func Infov(ctx context.Context, v any) { _ = "STUB: not implemented"; return }

func Infow(ctx context.Context, msg string, fields ...LogField) { _ = "STUB: not implemented"; return }

func Must(err error) { _ = "STUB: not implemented"; return }

func MustSetup(c logx.LogConf) { _ = "STUB: not implemented"; return }

func SetLevel(level uint32) { _ = "STUB: not implemented"; return }

func SetUp(c LogConf) error { _ = "STUB: not implemented"; return nil }

func Slow(ctx context.Context, v ...any) { _ = "STUB: not implemented"; return }

func Slowf(ctx context.Context, format string, v ...any) { _ = "STUB: not implemented"; return }

func Slowfn(ctx context.Context, fn func() any) { _ = "STUB: not implemented"; return }

func Slowv(ctx context.Context, v any) { _ = "STUB: not implemented"; return }

func Sloww(ctx context.Context, msg string, fields ...LogField) { _ = "STUB: not implemented"; return }

func getLogger(ctx context.Context) logx.Logger {
	_ = "STUB: not implemented"
	return *new(logx.Logger)
}
