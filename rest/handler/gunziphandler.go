package handler

import (
	"net/http"
)

const gzipEncoding = "gzip"

func GunzipHandler(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
