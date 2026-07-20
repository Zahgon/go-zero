package trace

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	gcodes "google.golang.org/grpc/codes"
)

const (
	GRPCStatusCodeKey = attribute.Key("rpc.grpc.status_code")

	RPCNameKey = attribute.Key("name")

	RPCMessageTypeKey = attribute.Key("message.type")

	RPCMessageIDKey = attribute.Key("message.id")

	RPCMessageCompressedSizeKey = attribute.Key("message.compressed_size")

	RPCMessageUncompressedSizeKey = attribute.Key("message.uncompressed_size")
)

var (
	RPCSystemGRPC = semconv.RPCSystemKey.String("grpc")

	RPCNameMessage = RPCNameKey.String("message")

	RPCMessageTypeSent = RPCMessageTypeKey.String("SENT")

	RPCMessageTypeReceived = RPCMessageTypeKey.String("RECEIVED")
)

func StatusCodeAttr(c gcodes.Code) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
