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

	Event interface {
		AskQuestion(ctx context.Context, in *EventReq, opts ...grpc.CallOption) (*EventResp, error)
	}

	defaultEvent struct {
		cli zrpc.Client
	}
)

func NewEvent(cli zrpc.Client) Event { _ = "STUB: not implemented"; return *new(Event) }

func (m *defaultEvent) AskQuestion(ctx context.Context, in *EventReq, opts ...grpc.CallOption) (*EventResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
