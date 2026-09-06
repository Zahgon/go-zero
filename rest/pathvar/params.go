package pathvar

import (
	"net/http"
)

var pathVars = contextKey("pathVars")

func Vars(r *http.Request) map[string]string { _ = "STUB: not implemented"; return nil }

func WithVars(r *http.Request, params map[string]string) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

type contextKey string

func (c contextKey) String() string { _ = "STUB: not implemented"; return "" }
