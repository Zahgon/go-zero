package logx

import "log"

type logWriter struct {
	logger *log.Logger
}

func newLogWriter(logger *log.Logger) logWriter { _ = "STUB: not implemented"; return *new(logWriter) }

func (lw logWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (lw logWriter) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
