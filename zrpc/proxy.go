package zrpc

import (
	"context"
	"sync"

	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/zrpc/internal"
	"google.golang.org/grpc"
)

type RpcProxy struct {
	backend      string
	clients      map[string]Client
	options      []internal.ClientOption
	singleFlight syncx.SingleFlight
	lock         sync.Mutex
}

func NewProxy(backend string, opts ...internal.ClientOption) *RpcProxy {
	_ = "STUB: not implemented"
	return nil
}

func (p *RpcProxy) TakeConn(ctx context.Context) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
