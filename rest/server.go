package rest

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/chain"
	"github.com/zeromicro/go-zero/rest/handler"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/internal"
)

type (
	RunOption func(*Server)

	StartOption = internal.StartOption

	Server struct {
		ngin   *engine
		router httpx.Router
	}
)

func MustNewServer(c RestConf, opts ...RunOption) *Server { _ = "STUB: not implemented"; return nil }

func NewServer(c RestConf, opts ...RunOption) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) AddRoute(r Route, opts ...RouteOption) { _ = "STUB: not implemented"; return }

func (s *Server) AddRoutes(rs []Route, opts ...RouteOption) { _ = "STUB: not implemented"; return }

func (s *Server) PrintRoutes() { _ = "STUB: not implemented"; return }

func (s *Server) Routes() []Route { _ = "STUB: not implemented"; return nil }

func (s *Server) Start() { _ = "STUB: not implemented"; return }

func (s *Server) StartWithOpts(opts ...StartOption) { _ = "STUB: not implemented"; return }

func (s *Server) Stop() { _ = "STUB: not implemented"; return }

func (s *Server) Use(middleware Middleware) { _ = "STUB: not implemented"; return }

func (s *Server) build() error { _ = "STUB: not implemented"; return nil }

func (s *Server) serve(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func ToMiddleware(handler func(next http.Handler) http.Handler) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

func WithChain(chn chain.Chain) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

func WithCors(origin ...string) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

func WithCorsHeaders(headers ...string) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

func WithCustomCors(middlewareFn func(header http.Header), notAllowedFn func(http.ResponseWriter),
	origin ...string) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

func WithFileServer(path string, fs http.FileSystem) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

func WithJwt(secret string) RouteOption { _ = "STUB: not implemented"; return *new(RouteOption) }

func WithJwtTransition(secret, prevSecret string) RouteOption {
	_ = "STUB: not implemented"
	return *new(RouteOption)
}

func WithMaxBytes(maxBytes int64) RouteOption { _ = "STUB: not implemented"; return *new(RouteOption) }

func WithMiddlewares(ms []Middleware, rs ...Route) []Route { _ = "STUB: not implemented"; return nil }

func WithMiddleware(middleware Middleware, rs ...Route) []Route {
	_ = "STUB: not implemented"
	return nil
}

func WithNotFoundHandler(handler http.Handler) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

func WithNotAllowedHandler(handler http.Handler) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

func WithPrefix(group string) RouteOption { _ = "STUB: not implemented"; return *new(RouteOption) }

func WithPriority() RouteOption { _ = "STUB: not implemented"; return *new(RouteOption) }

func WithRouter(router httpx.Router) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

func WithSignature(signature SignatureConf) RouteOption {
	_ = "STUB: not implemented"
	return *new(RouteOption)
}

func WithSSE() RouteOption { _ = "STUB: not implemented"; return *new(RouteOption) }

func WithTimeout(timeout time.Duration) RouteOption {
	_ = "STUB: not implemented"
	return *new(RouteOption)
}

func WithTLSConfig(cfg *tls.Config) RunOption { _ = "STUB: not implemented"; return *new(RunOption) }

func WithUnauthorizedCallback(callback handler.UnauthorizedCallback) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

func WithUnsignedCallback(callback handler.UnsignedCallback) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

func handleError(err error) { _ = "STUB: not implemented"; return }

func validateSecret(secret string) { _ = "STUB: not implemented"; return }

type corsRouter struct {
	httpx.Router
	middleware Middleware
}

func newCorsRouter(router httpx.Router, headerFn func(http.Header), origins ...string) httpx.Router {
	_ = "STUB: not implemented"
	return *new(httpx.Router)
}

func (c *corsRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type fileServingRouter struct {
	httpx.Router
	middleware Middleware
}

func newFileServingRouter(router httpx.Router, path string, fs http.FileSystem) httpx.Router {
	_ = "STUB: not implemented"
	return *new(httpx.Router)
}

func (f *fileServingRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
