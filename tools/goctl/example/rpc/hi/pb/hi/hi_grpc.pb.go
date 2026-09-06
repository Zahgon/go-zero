package hi

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion7

type GreetClient interface {
	SayHi(ctx context.Context, in *HiReq, opts ...grpc.CallOption) (*HiResp, error)
	SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
}

type greetClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetClient(cc grpc.ClientConnInterface) GreetClient {
	_ = "STUB: not implemented"
	return *new(GreetClient)
}

func (c *greetClient) SayHi(ctx context.Context, in *HiReq, opts ...grpc.CallOption) (*HiResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *greetClient) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GreetServer interface {
	SayHi(context.Context, *HiReq) (*HiResp, error)
	SayHello(context.Context, *HelloReq) (*HelloResp, error)
	mustEmbedUnimplementedGreetServer()
}

type UnimplementedGreetServer struct {
}

func (UnimplementedGreetServer) SayHi(context.Context, *HiReq) (*HiResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
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

func _Greet_SayHi_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func _Greet_SayHello_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

var Greet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hi.Greet",
	HandlerType: (*GreetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _Greet_SayHi_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _Greet_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hi.proto",
}

type EventClient interface {
	AskQuestion(ctx context.Context, in *EventReq, opts ...grpc.CallOption) (*EventResp, error)
}

type eventClient struct {
	cc grpc.ClientConnInterface
}

func NewEventClient(cc grpc.ClientConnInterface) EventClient {
	_ = "STUB: not implemented"
	return *new(EventClient)
}

func (c *eventClient) AskQuestion(ctx context.Context, in *EventReq, opts ...grpc.CallOption) (*EventResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type EventServer interface {
	AskQuestion(context.Context, *EventReq) (*EventResp, error)
	mustEmbedUnimplementedEventServer()
}

type UnimplementedEventServer struct {
}

func (UnimplementedEventServer) AskQuestion(context.Context, *EventReq) (*EventResp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UnimplementedEventServer) mustEmbedUnimplementedEventServer() {
	_ = "STUB: not implemented"
	return
}

type UnsafeEventServer interface {
	mustEmbedUnimplementedEventServer()
}

func RegisterEventServer(s grpc.ServiceRegistrar, srv EventServer) {
	_ = "STUB: not implemented"
	return
}

func _Event_AskQuestion_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

var Event_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hi.Event",
	HandlerType: (*EventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AskQuestion",
			Handler:    _Event_AskQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hi.proto",
}
