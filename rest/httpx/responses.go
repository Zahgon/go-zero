package httpx

import (
	"context"
	"io"
	"net/http"
	"sync"
)

var (
	errorHandler func(context.Context, error) (int, any)
	errorLock    sync.RWMutex
	okHandler    func(context.Context, any) any
	okLock       sync.RWMutex
)

func Error(w http.ResponseWriter, err error, fns ...func(w http.ResponseWriter, err error)) {
	_ = "STUB: not implemented"
	return
}

func ErrorCtx(ctx context.Context, w http.ResponseWriter, err error,
	fns ...func(w http.ResponseWriter, err error)) {
	_ = "STUB: not implemented"
	return
}

func Ok(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func OkJson(w http.ResponseWriter, v any) { _ = "STUB: not implemented"; return }

func OkJsonCtx(ctx context.Context, w http.ResponseWriter, v any) {
	_ = "STUB: not implemented"
	return
}

func SetErrorHandler(handler func(error) (int, any)) { _ = "STUB: not implemented"; return }

func SetErrorHandlerCtx(handlerCtx func(context.Context, error) (int, any)) {
	_ = "STUB: not implemented"
	return
}

func SetOkHandler(handler func(context.Context, any) any) { _ = "STUB: not implemented"; return }

func Stream(ctx context.Context, w http.ResponseWriter, fn func(w io.Writer) bool) {
	_ = "STUB: not implemented"
	return
}

func WriteJson(w http.ResponseWriter, code int, v any) { _ = "STUB: not implemented"; return }

func WriteJsonCtx(ctx context.Context, w http.ResponseWriter, code int, v any) {
	_ = "STUB: not implemented"
	return
}

func buildErrorHandler(ctx context.Context) func(error) (int, any) {
	_ = "STUB: not implemented"
	return nil
}

func doHandleError(w http.ResponseWriter, err error, handler func(error) (int, any),
	writeJson func(w http.ResponseWriter, code int, v any),
	fns ...func(w http.ResponseWriter, err error)) {
	_ = "STUB: not implemented"
	return
}

func doWriteJson(w http.ResponseWriter, code int, v any) error {
	_ = "STUB: not implemented"
	return nil
}
