package chain

import "net/http"

type (
	Chain interface {
		Append(middlewares ...Middleware) Chain
		Prepend(middlewares ...Middleware) Chain
		Then(h http.Handler) http.Handler
		ThenFunc(fn http.HandlerFunc) http.Handler
	}

	Middleware func(http.Handler) http.Handler

	chain struct {
		middlewares []Middleware
	}
)

func New(middlewares ...Middleware) Chain { _ = "STUB: not implemented"; return *new(Chain) }

func (c chain) Append(middlewares ...Middleware) Chain {
	_ = "STUB: not implemented"
	return *new(Chain)
}

func (c chain) Prepend(middlewares ...Middleware) Chain {
	_ = "STUB: not implemented"
	return *new(Chain)
}

func (c chain) Then(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (c chain) ThenFunc(fn http.HandlerFunc) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func join(a, b []Middleware) []Middleware { _ = "STUB: not implemented"; return nil }
