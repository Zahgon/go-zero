package gateway

import (
	"net/http"

	"github.com/fullstorydev/grpcurl"
	"github.com/golang/protobuf/jsonpb"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

const defaultHttpScheme = "http"

type (
	Server struct {
		*rest.Server
		upstreams     []Upstream
		conns         []zrpc.Client
		processHeader func(http.Header) []string
		dialer        func(conf zrpc.RpcClientConf) zrpc.Client
		middlewares   []rest.Middleware
	}

	Option func(svr *Server)
)

func MustNewServer(c GatewayConf, opts ...Option) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Start() { _ = "STUB: not implemented"; return }

func (s *Server) Stop() { _ = "STUB: not implemented"; return }

func (s *Server) build() error { _ = "STUB: not implemented"; return nil }

func (s *Server) buildChainHandler(handler http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func (s *Server) buildGrpcHandler(source grpcurl.DescriptorSource, resolver jsonpb.AnyResolver,
	cli zrpc.Client, rpcPath string) func(http.ResponseWriter, *http.Request) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) buildGrpcRoute(up Upstream, writer mr.Writer[rest.Route], cancel func(error)) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) buildHttpHandler(target *HttpClientConf) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func (s *Server) buildHttpRoute(up Upstream, writer mr.Writer[rest.Route]) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) ensureUpstreamNames() error { _ = "STUB: not implemented"; return nil }

func (s *Server) prepareMetadata(header http.Header) []string {
	_ = "STUB: not implemented"
	return nil
}

func WithHeaderProcessor(processHeader func(http.Header) []string) func(*Server) {
	_ = "STUB: not implemented"
	return nil
}

func WithMiddleware(middlewares ...rest.Middleware) func(*Server) {
	_ = "STUB: not implemented"
	return nil
}

func buildRequestWithNewTarget(r *http.Request, target *HttpClientConf) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createDescriptorSource(cli zrpc.Client, up Upstream) (grpcurl.DescriptorSource, error) {
	_ = "STUB: not implemented"
	return *new(grpcurl.DescriptorSource), nil
}

func WithDialer(dialer func(conf zrpc.RpcClientConf) zrpc.Client) func(*Server) {
	_ = "STUB: not implemented"
	return nil
}
