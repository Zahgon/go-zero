package logx

import (
	"context"
	"time"
)

func WithCallerSkip(skip int) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func WithContext(ctx context.Context) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func WithDuration(d time.Duration) Logger { _ = "STUB: not implemented"; return *new(Logger) }

type richLogger struct {
	ctx        context.Context
	callerSkip int
	fields     []LogField
}

func (l *richLogger) Debug(v ...any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Debugf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Debugfn(fn func() any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Debugv(v any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Debugw(msg string, fields ...LogField) { _ = "STUB: not implemented"; return }

func (l *richLogger) Error(v ...any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Errorf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Errorfn(fn func() any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Errorv(v any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Errorw(msg string, fields ...LogField) { _ = "STUB: not implemented"; return }

func (l *richLogger) Info(v ...any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Infof(format string, v ...any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Infofn(fn func() any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Infov(v any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Infow(msg string, fields ...LogField) { _ = "STUB: not implemented"; return }

func (l *richLogger) Slow(v ...any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Slowf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Slowfn(fn func() any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Slowv(v any) { _ = "STUB: not implemented"; return }

func (l *richLogger) Sloww(msg string, fields ...LogField) { _ = "STUB: not implemented"; return }

func (l *richLogger) WithCallerSkip(skip int) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func (l *richLogger) WithContext(ctx context.Context) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func (l *richLogger) WithDuration(duration time.Duration) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func (l *richLogger) WithFields(fields ...LogField) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func (l *richLogger) buildFields(fields ...LogField) []LogField {
	_ = "STUB: not implemented"
	return nil
}

func (l *richLogger) debug(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func (l *richLogger) err(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func (l *richLogger) info(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func (l *richLogger) slow(v any, fields ...LogField) { _ = "STUB: not implemented"; return }
