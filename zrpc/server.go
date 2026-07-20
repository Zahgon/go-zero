package zrpc

import (
	"time"

	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/zrpc/internal"
	"google.golang.org/grpc"
)

type RpcServer struct {
	server   internal.Server
	register internal.RegisterFn
}

func MustNewServer(c RpcServerConf, register internal.RegisterFn) *RpcServer {
	_ = "STUB: not implemented"
	return nil
}

func NewServer(c RpcServerConf, register internal.RegisterFn) (*RpcServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rs *RpcServer) AddOptions(options ...grpc.ServerOption) { _ = "STUB: not implemented"; return }

func (rs *RpcServer) AddStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) {
	_ = "STUB: not implemented"
	return
}

func (rs *RpcServer) AddUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) {
	_ = "STUB: not implemented"
	return
}

func (rs *RpcServer) Start() { _ = "STUB: not implemented"; return }

func (rs *RpcServer) Stop() { _ = "STUB: not implemented"; return }

func DontLogContentForMethod(method string) { _ = "STUB: not implemented"; return }

func SetServerSlowThreshold(threshold time.Duration) { _ = "STUB: not implemented"; return }

func setupAuthInterceptors(svr internal.Server, c RpcServerConf) error {
	_ = "STUB: not implemented"
	return nil
}

func setupStreamInterceptors(svr internal.Server, c RpcServerConf) {
	_ = "STUB: not implemented"
	return
}

func setupUnaryInterceptors(svr internal.Server, c RpcServerConf, metrics *stat.Metrics) {
	_ = "STUB: not implemented"
	return
}
