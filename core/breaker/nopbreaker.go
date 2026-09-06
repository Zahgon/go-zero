package breaker

import "context"

const nopBreakerName = "nopBreaker"

type nopBreaker struct{}

func NopBreaker() Breaker { _ = "STUB: not implemented"; return *new(Breaker) }

func (b nopBreaker) Name() string { _ = "STUB: not implemented"; return "" }

func (b nopBreaker) Allow() (Promise, error) { _ = "STUB: not implemented"; return *new(Promise), nil }

func (b nopBreaker) AllowCtx(_ context.Context) (Promise, error) {
	_ = "STUB: not implemented"
	return *new(Promise), nil
}

func (b nopBreaker) Do(req func() error) error { _ = "STUB: not implemented"; return nil }

func (b nopBreaker) DoCtx(_ context.Context, req func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (b nopBreaker) DoWithAcceptable(req func() error, _ Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func (b nopBreaker) DoWithAcceptableCtx(_ context.Context, req func() error, _ Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func (b nopBreaker) DoWithFallback(req func() error, _ Fallback) error {
	_ = "STUB: not implemented"
	return nil
}

func (b nopBreaker) DoWithFallbackCtx(_ context.Context, req func() error, _ Fallback) error {
	_ = "STUB: not implemented"
	return nil
}

func (b nopBreaker) DoWithFallbackAcceptable(req func() error, _ Fallback, _ Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func (b nopBreaker) DoWithFallbackAcceptableCtx(_ context.Context, req func() error,
	_ Fallback, _ Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

type nopPromise struct{}

func (p nopPromise) Accept() { _ = "STUB: not implemented"; return }

func (p nopPromise) Reject(_ string) { _ = "STUB: not implemented"; return }
