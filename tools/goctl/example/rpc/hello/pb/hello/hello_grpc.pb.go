package hello

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion7

type GreetClient interface {
	SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
}

type greetClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetClient(cc grpc.ClientConnInterface) GreetClient {
	_ = "STUB: not implemented"
	return *new(GreetClient)
}

func (c *greetClient) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GreetServer interface {
	SayHello(context.Context, *HelloReq) (*HelloResp, error)
	mustEmbedUnimplementedGreetServer()
}

type UnimplementedGreetServer struct {
}

func (UnimplementedGreetServer) SayHello(context.Context, *HelloReq) (*HelloResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedGreetServer) mustEmbedUnimplementedGreetServer() {
	_ = "STUB: not implemented"
	return
}

type UnsafeGreetServer interface {
	mustEmbedUnimplementedGreetServer()
}

func RegisterGreetServer(s grpc.ServiceRegistrar, srv GreetServer) {
	_ = "STUB: not implemented"
	return
}

func _Greet_SayHello_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

var Greet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Greet",
	HandlerType: (*GreetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greet_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
