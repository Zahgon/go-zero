package zrpc

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/zrpc/internal"
	"google.golang.org/grpc"
)

var (
	WithDialOption = internal.WithDialOption

	WithNonBlock = internal.WithNonBlock

	WithBlock = internal.WithBlock

	WithStreamClientInterceptor = internal.WithStreamClientInterceptor

	WithTimeout = internal.WithTimeout

	WithTransportCredentials = internal.WithTransportCredentials

	WithUnaryClientInterceptor = internal.WithUnaryClientInterceptor
)

type (
	Client = internal.Client

	ClientOption = internal.ClientOption

	RpcClient struct {
		client Client
	}
)

func MustNewClient(c RpcClientConf, options ...ClientOption) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

func NewClient(c RpcClientConf, options ...ClientOption) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func NewClientWithTarget(target string, opts ...ClientOption) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func (rc *RpcClient) Conn() *grpc.ClientConn { _ = "STUB: not implemented"; return nil }

func DontLogClientContentForMethod(method string) { _ = "STUB: not implemented"; return }

func SetClientSlowThreshold(threshold time.Duration) { _ = "STUB: not implemented"; return }

func SetHashKey(ctx context.Context, key string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func WithCallTimeout(timeout time.Duration) grpc.CallOption {
	_ = "STUB: not implemented"
	return *new(grpc.CallOption)
}

func makeLBServiceConfig(balancerName string) string { _ = "STUB: not implemented"; return "" }
