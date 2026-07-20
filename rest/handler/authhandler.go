package handler

import (
	"errors"
	"net/http"
)

const (
	jwtAudience    = "aud"
	jwtExpire      = "exp"
	jwtId          = "jti"
	jwtIssueAt     = "iat"
	jwtIssuer      = "iss"
	jwtNotBefore   = "nbf"
	jwtSubject     = "sub"
	noDetailReason = "no detail reason"
)

var (
	errInvalidToken = errors.New("invalid auth token")
	errNoClaims     = errors.New("no auth params")
)

type (
	AuthorizeOptions struct {
		PrevSecret string
		Callback   UnauthorizedCallback
	}

	UnauthorizedCallback func(w http.ResponseWriter, r *http.Request, err error)

	AuthorizeOption func(opts *AuthorizeOptions)
)

func Authorize(secret string, opts ...AuthorizeOption) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func WithPrevSecret(secret string) AuthorizeOption {
	_ = "STUB: not implemented"
	return *new(AuthorizeOption)
}

func WithUnauthorizedCallback(callback UnauthorizedCallback) AuthorizeOption {
	_ = "STUB: not implemented"
	return *new(AuthorizeOption)
}

func detailAuthLog(r *http.Request, reason string) { _ = "STUB: not implemented"; return }

func unauthorized(w http.ResponseWriter, r *http.Request, err error, callback UnauthorizedCallback) {
	_ = "STUB: not implemented"
	return
}
