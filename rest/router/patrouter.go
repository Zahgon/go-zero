package router

import (
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/core/search"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	allowHeader          = "Allow"
	allowMethodSeparator = ", "
)

var (
	ErrInvalidMethod = errors.New("not a valid http method")

	ErrInvalidPath = errors.New("path must begin with '/'")
)

type patRouter struct {
	trees      map[string]*search.Tree
	notFound   http.Handler
	notAllowed http.Handler
}

func NewRouter() httpx.Router { _ = "STUB: not implemented"; return *new(httpx.Router) }

func (pr *patRouter) Handle(method, reqPath string, handler http.Handler) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *patRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (pr *patRouter) SetNotFoundHandler(handler http.Handler) { _ = "STUB: not implemented"; return }

func (pr *patRouter) SetNotAllowedHandler(handler http.Handler) { _ = "STUB: not implemented"; return }

func (pr *patRouter) handleNotFound(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (pr *patRouter) methodsAllowed(method, path string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func validMethod(method string) bool { _ = "STUB: not implemented"; return false }
