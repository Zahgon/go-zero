package trace

import (
	"context"
	"sync"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	kindZipkin   = "zipkin"
	kindOtlpGrpc = "otlpgrpc"
	kindOtlpHttp = "otlphttp"
	kindFile     = "file"
)

var (
	once           sync.Once
	tp             *sdktrace.TracerProvider
	shutdownOnceFn = sync.OnceFunc(func() {
		if tp != nil {
			_ = tp.Shutdown(context.Background())
		}
	})
)

func StartAgent(c Config) { _ = "STUB: not implemented"; return }

func StopAgent() { _ = "STUB: not implemented"; return }

func createExporter(c Config) (sdktrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanExporter), nil
}

func startAgent(c Config) error { _ = "STUB: not implemented"; return nil }
