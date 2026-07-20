package trace

import (
	"context"

	ztrace "github.com/zeromicro/go-zero/internal/trace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const localhost = "127.0.0.1"

var (
	SpanIDFromContext = ztrace.SpanIDFromContext

	TraceIDFromContext = ztrace.TraceIDFromContext
)

func ParseFullMethod(fullMethod string) (string, []attribute.KeyValue) {
	_ = "STUB: not implemented"
	return "", nil
}

func PeerAttr(addr string) []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func PeerFromCtx(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func SpanInfo(fullMethod, peerAddress string) (string, []attribute.KeyValue) {
	_ = "STUB: not implemented"
	return "", nil
}

func TracerFromContext(ctx context.Context) (tracer trace.Tracer) {
	_ = "STUB: not implemented"
	return *new(trace.Tracer)
}
