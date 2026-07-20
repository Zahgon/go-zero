package breaker

import (
	"context"
	"sync"
)

var (
	lock     sync.RWMutex
	breakers = make(map[string]Breaker)
)

func Do(name string, req func() error) error { _ = "STUB: not implemented"; return nil }

func DoCtx(ctx context.Context, name string, req func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func DoWithAcceptable(name string, req func() error, acceptable Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func DoWithAcceptableCtx(ctx context.Context, name string, req func() error,
	acceptable Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func DoWithFallback(name string, req func() error, fallback Fallback) error {
	_ = "STUB: not implemented"
	return nil
}

func DoWithFallbackCtx(ctx context.Context, name string, req func() error, fallback Fallback) error {
	_ = "STUB: not implemented"
	return nil
}

func DoWithFallbackAcceptable(name string, req func() error, fallback Fallback,
	acceptable Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func DoWithFallbackAcceptableCtx(ctx context.Context, name string, req func() error,
	fallback Fallback, acceptable Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func GetBreaker(name string) Breaker { _ = "STUB: not implemented"; return *new(Breaker) }

func NoBreakerFor(name string) { _ = "STUB: not implemented"; return }

func do(name string, execute func(b Breaker) error) error { _ = "STUB: not implemented"; return nil }
