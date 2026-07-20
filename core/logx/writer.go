package logx

import (
	"io"
	"sync"
)

type (
	Writer interface {
		Alert(v any)

		Close() error

		Debug(v any, fields ...LogField)

		Error(v any, fields ...LogField)

		Info(v any, fields ...LogField)

		Severe(v any)

		Slow(v any, fields ...LogField)

		Stack(v any)

		Stat(v any, fields ...LogField)
	}

	atomicWriter struct {
		writer Writer
		lock   sync.RWMutex
	}

	comboWriter struct {
		writers []Writer
	}

	concreteWriter struct {
		infoLog   io.WriteCloser
		errorLog  io.WriteCloser
		severeLog io.WriteCloser
		slowLog   io.WriteCloser
		statLog   io.WriteCloser
		stackLog  io.Writer
	}
)

func NewWriter(w io.Writer) Writer { _ = "STUB: not implemented"; return *new(Writer) }

func (w *atomicWriter) Load() Writer { _ = "STUB: not implemented"; return *new(Writer) }

func (w *atomicWriter) Store(v Writer) { _ = "STUB: not implemented"; return }

func (w *atomicWriter) StoreIfNil(v Writer) Writer { _ = "STUB: not implemented"; return *new(Writer) }

func (w *atomicWriter) Swap(v Writer) Writer { _ = "STUB: not implemented"; return *new(Writer) }

func (c comboWriter) Alert(v any) { _ = "STUB: not implemented"; return }

func (c comboWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (c comboWriter) Debug(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func (c comboWriter) Error(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func (c comboWriter) Info(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func (c comboWriter) Severe(v any) { _ = "STUB: not implemented"; return }

func (c comboWriter) Slow(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func (c comboWriter) Stack(v any) { _ = "STUB: not implemented"; return }

func (c comboWriter) Stat(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func newConsoleWriter() Writer { _ = "STUB: not implemented"; return *new(Writer) }

func newFileWriter(c LogConf) (Writer, error) { _ = "STUB: not implemented"; return *new(Writer), nil }

func (w *concreteWriter) Alert(v any) { _ = "STUB: not implemented"; return }

func (w *concreteWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (w *concreteWriter) Debug(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func (w *concreteWriter) Error(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func (w *concreteWriter) Info(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func (w *concreteWriter) Severe(v any) { _ = "STUB: not implemented"; return }

func (w *concreteWriter) Slow(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

func (w *concreteWriter) Stack(v any) { _ = "STUB: not implemented"; return }

func (w *concreteWriter) Stat(v any, fields ...LogField) { _ = "STUB: not implemented"; return }

type nopWriter struct{}

func (n nopWriter) Alert(_ any) { _ = "STUB: not implemented"; return }

func (n nopWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (n nopWriter) Debug(_ any, _ ...LogField) { _ = "STUB: not implemented"; return }

func (n nopWriter) Error(_ any, _ ...LogField) { _ = "STUB: not implemented"; return }

func (n nopWriter) Info(_ any, _ ...LogField) { _ = "STUB: not implemented"; return }

func (n nopWriter) Severe(_ any) { _ = "STUB: not implemented"; return }

func (n nopWriter) Slow(_ any, _ ...LogField) { _ = "STUB: not implemented"; return }

func (n nopWriter) Stack(_ any) { _ = "STUB: not implemented"; return }

func (n nopWriter) Stat(_ any, _ ...LogField) { _ = "STUB: not implemented"; return }

func buildPlainFields(fields logEntry) []string { _ = "STUB: not implemented"; return nil }

func marshalJson(t interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func mergeGlobalFields(fields []LogField) []LogField { _ = "STUB: not implemented"; return nil }

func output(writer io.Writer, level string, val any, fields ...LogField) {
	_ = "STUB: not implemented"
	return
}

func processFieldValue(value any) any { _ = "STUB: not implemented"; return *new(any) }

func wrapLevelWithColor(level string) string { _ = "STUB: not implemented"; return "" }

func writeJson(writer io.Writer, info any) { _ = "STUB: not implemented"; return }

func writePlainAny(writer io.Writer, level string, val any, fields ...string) {
	_ = "STUB: not implemented"
	return
}

func writePlainText(writer io.Writer, level, msg string, fields ...string) {
	_ = "STUB: not implemented"
	return
}

func writePlainValue(writer io.Writer, level string, val any, fields ...string) {
	_ = "STUB: not implemented"
	return
}
