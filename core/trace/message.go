package trace

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
)

const messageEvent = "message"

var (
	MessageSent = messageType(RPCMessageTypeSent)

	MessageReceived = messageType(RPCMessageTypeReceived)
)

type messageType attribute.KeyValue

func (m messageType) Event(ctx context.Context, id int, message any) {
	_ = "STUB: not implemented"
	return
}
