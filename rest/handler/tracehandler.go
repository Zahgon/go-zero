package handler

import (
	"net/http"
)

type (
	TraceOption func(options *traceOptions)

	traceOptions struct {
		traceIgnorePaths []string
	}
)

func TraceHandler(serviceName, path string, opts ...TraceOption) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func WithTraceIgnorePaths(traceIgnorePaths []string) TraceOption {
	_ = "STUB: not implemented"
	return *new(TraceOption)
}
