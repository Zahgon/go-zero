package client

import (
	"context"

	"github.com/zeromicro/go-zero/tools/goctl/example/rpc/hi/pb/hi"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	EventReq  = hi.EventReq
	EventResp = hi.EventResp
	HelloReq  = hi.HelloReq
	HelloResp = hi.HelloResp
	HiReq     = hi.HiReq
	HiResp    = hi.HiResp

	Greet interface {
		SayHi(ctx context.Context, in *HiReq, opts ...grpc.CallOption) (*HiResp, error)
		SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
	}

	defaultGreet struct {
		cli zrpc.Client
	}
)

func NewGreet(cli zrpc.Client) Greet { _ = "STUB: not implemented"; return *new(Greet) }

func (m *defaultGreet) SayHi(ctx context.Context, in *HiReq, opts ...grpc.CallOption) (*HiResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *defaultGreet) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
