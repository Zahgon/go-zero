package rest

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/load"
	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/rest/chain"
	"github.com/zeromicro/go-zero/rest/handler"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/internal"
)

const topCpuUsage = 1000

var ErrSignatureConfig = errors.New("bad config for Signature")

type engine struct {
	conf   RestConf
	routes []featuredRoutes

	timeout              time.Duration
	unauthorizedCallback handler.UnauthorizedCallback
	unsignedCallback     handler.UnsignedCallback
	chain                chain.Chain
	middlewares          []Middleware
	shedder              load.Shedder
	priorityShedder      load.Shedder
	tlsConfig            *tls.Config
}

func newEngine(c RestConf) *engine { _ = "STUB: not implemented"; return nil }

func (ng *engine) addRoutes(r featuredRoutes) { _ = "STUB: not implemented"; return }

func (ng *engine) appendAuthHandler(fr featuredRoutes, chn chain.Chain,
	verifier func(chain.Chain) chain.Chain) chain.Chain {
	_ = "STUB: not implemented"
	return *new(chain.Chain)
}

func (ng *engine) bindFeaturedRoutes(router httpx.Router, fr featuredRoutes, metrics *stat.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (ng *engine) bindRoute(fr featuredRoutes, router httpx.Router, metrics *stat.Metrics,
	route Route, verifier func(chain.Chain) chain.Chain) error {
	_ = "STUB: not implemented"
	return nil
}

func (ng *engine) bindRoutes(router httpx.Router) error { _ = "STUB: not implemented"; return nil }

func (ng *engine) buildChainWithNativeMiddlewares(fr featuredRoutes, route Route,
	metrics *stat.Metrics) chain.Chain {
	_ = "STUB: not implemented"
	return *new(chain.Chain)
}

func (ng *engine) checkedMaxBytes(bytes int64) int64 { _ = "STUB: not implemented"; return 0 }

func (ng *engine) checkedTimeout(timeout *time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (ng *engine) createMetrics() *stat.Metrics { _ = "STUB: not implemented"; return nil }

func (ng *engine) getLogHandler() func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func (ng *engine) getShedder(priority bool) load.Shedder {
	_ = "STUB: not implemented"
	return *new(load.Shedder)
}

func (ng *engine) hasTimeout() bool { _ = "STUB: not implemented"; return false }

func (ng *engine) mightUpdateTimeout(r featuredRoutes) { _ = "STUB: not implemented"; return }

func (ng *engine) notFoundHandler(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (ng *engine) print() { _ = "STUB: not implemented"; return }

func (ng *engine) setTlsConfig(cfg *tls.Config) { _ = "STUB: not implemented"; return }

func (ng *engine) setUnauthorizedCallback(callback handler.UnauthorizedCallback) {
	_ = "STUB: not implemented"
	return
}

func (ng *engine) setUnsignedCallback(callback handler.UnsignedCallback) {
	_ = "STUB: not implemented"
	return
}

func (ng *engine) signatureVerifier(signature signatureSetting) (func(chain.Chain) chain.Chain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ng *engine) start(router httpx.Router, opts ...StartOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (ng *engine) use(middleware Middleware) { _ = "STUB: not implemented"; return }

func (ng *engine) withNetworkTimeout() internal.StartOption {
	_ = "STUB: not implemented"
	return *new(internal.StartOption)
}

func buildSSERoutes(routes []Route) []Route { _ = "STUB: not implemented"; return nil }

func convertMiddleware(ware Middleware) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
