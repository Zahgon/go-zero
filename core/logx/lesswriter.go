package logx

import "io"

type lessWriter struct {
	*limitedExecutor
	writer io.Writer
}

func newLessWriter(writer io.Writer, milliseconds int) *lessWriter {
	_ = "STUB: not implemented"
	return nil
}

func (w *lessWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
