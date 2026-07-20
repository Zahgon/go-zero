package trace

import (
	"context"

	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

var _ propagation.TextMapCarrier = (*metadataSupplier)(nil)

type metadataSupplier struct {
	metadata *metadata.MD
}

func (s *metadataSupplier) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (s *metadataSupplier) Set(key, value string) { _ = "STUB: not implemented"; return }

func (s *metadataSupplier) Keys() []string { _ = "STUB: not implemented"; return nil }

func Inject(ctx context.Context, p propagation.TextMapPropagator, metadata *metadata.MD) {
	_ = "STUB: not implemented"
	return
}

func Extract(ctx context.Context, p propagation.TextMapPropagator, metadata *metadata.MD) (
	baggage.Baggage, sdktrace.SpanContext) {
	_ = "STUB: not implemented"
	return *new(baggage.Baggage), *new(sdktrace.SpanContext)
}
