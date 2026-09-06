package sqlx

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var sqlAttributeKey = attribute.Key("sql.method")

func startSpan(ctx context.Context, method string) (context.Context, oteltrace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(oteltrace.Span)
}

func endSpan(span oteltrace.Span, err error) { _ = "STUB: not implemented"; return }
