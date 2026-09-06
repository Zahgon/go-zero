package generator

import (
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
)

type rpcTypeRef struct {
	GoRef string

	ImportPath string
}

func resolveRPCTypeRef(protoType, mainPbPackage, mainGoPackage string, pkgMap map[string]parser.ImportedProto) rpcTypeRef {
	_ = "STUB: not implemented"
	return *new(rpcTypeRef)
}

func resolveCallTypeRef(protoType, mainPbPackage, mainGoPackage string, pkgMap map[string]parser.ImportedProto) (typeName, aliasEntry, importPath string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

var googleWKTTable = map[string]rpcTypeRef{
	"Empty":       {GoRef: "emptypb.Empty", ImportPath: "google.golang.org/protobuf/types/known/emptypb"},
	"Timestamp":   {GoRef: "timestamppb.Timestamp", ImportPath: "google.golang.org/protobuf/types/known/timestamppb"},
	"Duration":    {GoRef: "durationpb.Duration", ImportPath: "google.golang.org/protobuf/types/known/durationpb"},
	"Any":         {GoRef: "anypb.Any", ImportPath: "google.golang.org/protobuf/types/known/anypb"},
	"StringValue": {GoRef: "wrapperspb.StringValue", ImportPath: "google.golang.org/protobuf/types/known/wrapperspb"},
	"Int32Value":  {GoRef: "wrapperspb.Int32Value", ImportPath: "google.golang.org/protobuf/types/known/wrapperspb"},
	"Int64Value":  {GoRef: "wrapperspb.Int64Value", ImportPath: "google.golang.org/protobuf/types/known/wrapperspb"},
	"BoolValue":   {GoRef: "wrapperspb.BoolValue", ImportPath: "google.golang.org/protobuf/types/known/wrapperspb"},
	"BytesValue":  {GoRef: "wrapperspb.BytesValue", ImportPath: "google.golang.org/protobuf/types/known/wrapperspb"},
	"FloatValue":  {GoRef: "wrapperspb.FloatValue", ImportPath: "google.golang.org/protobuf/types/known/wrapperspb"},
	"DoubleValue": {GoRef: "wrapperspb.DoubleValue", ImportPath: "google.golang.org/protobuf/types/known/wrapperspb"},
	"UInt32Value": {GoRef: "wrapperspb.UInt32Value", ImportPath: "google.golang.org/protobuf/types/known/wrapperspb"},
	"UInt64Value": {GoRef: "wrapperspb.UInt64Value", ImportPath: "google.golang.org/protobuf/types/known/wrapperspb"},
	"Struct":      {GoRef: "structpb.Struct", ImportPath: "google.golang.org/protobuf/types/known/structpb"},
	"Value":       {GoRef: "structpb.Value", ImportPath: "google.golang.org/protobuf/types/known/structpb"},
	"ListValue":   {GoRef: "structpb.ListValue", ImportPath: "google.golang.org/protobuf/types/known/structpb"},
	"FieldMask":   {GoRef: "fieldmaskpb.FieldMask", ImportPath: "google.golang.org/protobuf/types/known/fieldmaskpb"},
}

func resolveGoogleWKT(typeName string) rpcTypeRef {
	_ = "STUB: not implemented"
	return *new(rpcTypeRef)
}
