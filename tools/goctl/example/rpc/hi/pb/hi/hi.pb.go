package hi

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HiReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	In string `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
}

func (x *HiReq) Reset() { _ = "STUB: not implemented"; return }

func (x *HiReq) String() string { _ = "STUB: not implemented"; return "" }

func (*HiReq) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HiReq) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HiReq) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HiReq) GetIn() string { _ = "STUB: not implemented"; return "" }

type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	In string `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
}

func (x *HelloReq) Reset() { _ = "STUB: not implemented"; return }

func (x *HelloReq) String() string { _ = "STUB: not implemented"; return "" }

func (*HelloReq) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HelloReq) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HelloReq) GetIn() string { _ = "STUB: not implemented"; return "" }

type HiResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HiResp) Reset() { _ = "STUB: not implemented"; return }

func (x *HiResp) String() string { _ = "STUB: not implemented"; return "" }

func (*HiResp) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HiResp) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HiResp) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HiResp) GetMsg() string { _ = "STUB: not implemented"; return "" }

type HelloResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloResp) Reset() { _ = "STUB: not implemented"; return }

func (x *HelloResp) String() string { _ = "STUB: not implemented"; return "" }

func (*HelloResp) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *HelloResp) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*HelloResp) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *HelloResp) GetMsg() string { _ = "STUB: not implemented"; return "" }

type EventReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventReq) Reset() { _ = "STUB: not implemented"; return }

func (x *EventReq) String() string { _ = "STUB: not implemented"; return "" }

func (*EventReq) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *EventReq) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*EventReq) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type EventResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventResp) Reset() { _ = "STUB: not implemented"; return }

func (x *EventResp) String() string { _ = "STUB: not implemented"; return "" }

func (*EventResp) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *EventResp) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*EventResp) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

var File_hi_proto protoreflect.FileDescriptor

var file_hi_proto_rawDesc = []byte{
	0x0a, 0x08, 0x68, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x68, 0x69, 0x22, 0x17,
	0x0a, 0x05, 0x48, 0x69, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x6e, 0x22, 0x1a, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x6e, 0x22, 0x1a, 0x0a, 0x06, 0x48, 0x69, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x1d, 0x0a, 0x09, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x0a,
	0x0a, 0x08, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x22, 0x0b, 0x0a, 0x09, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x32, 0x50, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x12, 0x1e, 0x0a, 0x05, 0x53, 0x61, 0x79, 0x48, 0x69, 0x12, 0x09, 0x2e, 0x68, 0x69, 0x2e, 0x48,
	0x69, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x68, 0x69, 0x2e, 0x48, 0x69, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x27, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0c, 0x2e, 0x68,
	0x69, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x68, 0x69, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x32, 0x33, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x0b, 0x41, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0c, 0x2e, 0x68, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x0d, 0x2e, 0x68, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x68, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hi_proto_rawDescOnce sync.Once
	file_hi_proto_rawDescData = file_hi_proto_rawDesc
)

func file_hi_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_hi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_hi_proto_goTypes = []any{
	(*HiReq)(nil),
	(*HelloReq)(nil),
	(*HiResp)(nil),
	(*HelloResp)(nil),
	(*EventReq)(nil),
	(*EventResp)(nil),
}
var file_hi_proto_depIdxs = []int32{
	0,
	1,
	4,
	2,
	3,
	5,
	3,
	0,
	0,
	0,
	0,
}

func init()               { file_hi_proto_init() }
func file_hi_proto_init() { _ = "STUB: not implemented"; return }
