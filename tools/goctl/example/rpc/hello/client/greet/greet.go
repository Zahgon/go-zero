package client

import (
	"context"

	"github.com/zeromicro/go-zero/tools/goctl/example/rpc/hello/pb/hello"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	HelloReq  = hello.HelloReq
	HelloResp = hello.HelloResp

	Greet interface {
		SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
	}

	defaultGreet struct {
		cli zrpc.Client
	}
)

func NewGreet(cli zrpc.Client) Greet { _ = "STUB: not implemented"; return *new(Greet) }

func (m *defaultGreet) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
