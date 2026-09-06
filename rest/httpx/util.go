package httpx

import (
	"net/http"
)

const (
	xForwardedFor = "X-Forwarded-For"
	arraySuffix   = "[]"

	maxFormParamCount = 2048
)

func GetFormValues(r *http.Request) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetRemoteAddr(r *http.Request) string { _ = "STUB: not implemented"; return "" }
