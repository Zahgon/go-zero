package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/codec"
)

const contentSecurity = "X-Content-Security"

type UnsignedCallback func(w http.ResponseWriter, r *http.Request, next http.Handler, strict bool, code int)

func ContentSecurityHandler(decrypters map[string]codec.RsaDecrypter, tolerance time.Duration,
	strict bool, callbacks ...UnsignedCallback) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func LimitContentSecurityHandler(limitBytes int64, decrypters map[string]codec.RsaDecrypter,
	tolerance time.Duration, strict bool, callbacks ...UnsignedCallback) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func executeCallbacks(w http.ResponseWriter, r *http.Request, next http.Handler, strict bool,
	code int, callbacks []UnsignedCallback) {
	_ = "STUB: not implemented"
	return
}

func handleVerificationFailure(w http.ResponseWriter, r *http.Request, next http.Handler,
	strict bool, _ int) {
	_ = "STUB: not implemented"
	return
}
