package breaker

import (
	"context"
	"errors"
	"sync"
)

const numHistoryReasons = 5

var ErrServiceUnavailable = errors.New("circuit breaker is open")

type (
	Acceptable func(err error) bool

	Breaker interface {
		Name() string

		Allow() (Promise, error)

		AllowCtx(ctx context.Context) (Promise, error)

		Do(req func() error) error

		DoCtx(ctx context.Context, req func() error) error

		DoWithAcceptable(req func() error, acceptable Acceptable) error

		DoWithAcceptableCtx(ctx context.Context, req func() error, acceptable Acceptable) error

		DoWithFallback(req func() error, fallback Fallback) error

		DoWithFallbackCtx(ctx context.Context, req func() error, fallback Fallback) error

		DoWithFallbackAcceptable(req func() error, fallback Fallback, acceptable Acceptable) error

		DoWithFallbackAcceptableCtx(ctx context.Context, req func() error, fallback Fallback,
			acceptable Acceptable) error
	}

	Fallback func(err error) error

	Option func(breaker *circuitBreaker)

	Promise interface {
		Accept()

		Reject(reason string)
	}

	internalPromise interface {
		Accept()
		Reject()
	}

	circuitBreaker struct {
		name string
		throttle
	}

	internalThrottle interface {
		allow() (internalPromise, error)
		doReq(req func() error, fallback Fallback, acceptable Acceptable) error
	}

	throttle interface {
		allow() (Promise, error)
		doReq(req func() error, fallback Fallback, acceptable Acceptable) error
	}
)

func NewBreaker(opts ...Option) Breaker { _ = "STUB: not implemented"; return *new(Breaker) }

func (cb *circuitBreaker) Allow() (Promise, error) {
	_ = "STUB: not implemented"
	return *new(Promise), nil
}

func (cb *circuitBreaker) AllowCtx(ctx context.Context) (Promise, error) {
	_ = "STUB: not implemented"
	return *new(Promise), nil
}

func (cb *circuitBreaker) Do(req func() error) error { _ = "STUB: not implemented"; return nil }

func (cb *circuitBreaker) DoCtx(ctx context.Context, req func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *circuitBreaker) DoWithAcceptable(req func() error, acceptable Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *circuitBreaker) DoWithAcceptableCtx(ctx context.Context, req func() error,
	acceptable Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *circuitBreaker) DoWithFallback(req func() error, fallback Fallback) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *circuitBreaker) DoWithFallbackCtx(ctx context.Context, req func() error,
	fallback Fallback) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *circuitBreaker) DoWithFallbackAcceptable(req func() error, fallback Fallback,
	acceptable Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *circuitBreaker) DoWithFallbackAcceptableCtx(ctx context.Context, req func() error,
	fallback Fallback, acceptable Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *circuitBreaker) Name() string { _ = "STUB: not implemented"; return "" }

func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

func defaultAcceptable(err error) bool { _ = "STUB: not implemented"; return false }

type loggedThrottle struct {
	name string
	internalThrottle
	errWin *errorWindow
}

func newLoggedThrottle(name string, t internalThrottle) loggedThrottle {
	_ = "STUB: not implemented"
	return *new(loggedThrottle)
}

func (lt loggedThrottle) allow() (Promise, error) {
	_ = "STUB: not implemented"
	return *new(Promise), nil
}

func (lt loggedThrottle) doReq(req func() error, fallback Fallback, acceptable Acceptable) error {
	_ = "STUB: not implemented"
	return nil
}

func (lt loggedThrottle) logError(err error) error { _ = "STUB: not implemented"; return nil }

type errorWindow struct {
	reasons [numHistoryReasons]string
	index   int
	count   int
	lock    sync.Mutex
}

func (ew *errorWindow) add(reason string) { _ = "STUB: not implemented"; return }

func (ew *errorWindow) String() string { _ = "STUB: not implemented"; return "" }

type promiseWithReason struct {
	promise internalPromise
	errWin  *errorWindow
}

func (p promiseWithReason) Accept() { _ = "STUB: not implemented"; return }

func (p promiseWithReason) Reject(reason string) { _ = "STUB: not implemented"; return }
