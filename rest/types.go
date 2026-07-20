package rest

import (
	"net/http"
	"time"
)

type (
	Middleware func(next http.HandlerFunc) http.HandlerFunc

	Route struct {
		Method  string
		Path    string
		Handler http.HandlerFunc
	}

	RouteOption func(r *featuredRoutes)

	jwtSetting struct {
		enabled    bool
		secret     string
		prevSecret string
	}

	signatureSetting struct {
		SignatureConf
		enabled bool
	}

	featuredRoutes struct {
		timeout   *time.Duration
		priority  bool
		jwt       jwtSetting
		signature signatureSetting
		sse       bool
		routes    []Route
		maxBytes  int64
	}
)
