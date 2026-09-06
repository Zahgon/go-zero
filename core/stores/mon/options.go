package mon

import (
	"reflect"
	"time"

	"github.com/zeromicro/go-zero/core/syncx"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const defaultTimeout = time.Second * 3

var (
	slowThreshold = syncx.ForAtomicDuration(defaultSlowThreshold)
	logMon        = syncx.ForAtomicBool(true)
	logSlowMon    = syncx.ForAtomicBool(true)
)

type (
	Option func(opts *clientOptions)

	TypeCodec struct {
		ValueType reflect.Type
		Encoder   bson.ValueEncoder
		Decoder   bson.ValueDecoder
	}

	clientOptions = options.ClientOptions
)

func DisableLog() { _ = "STUB: not implemented"; return }

func DisableInfoLog() { _ = "STUB: not implemented"; return }

func SetSlowThreshold(threshold time.Duration) { _ = "STUB: not implemented"; return }

func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTypeCodec(typeCodecs ...TypeCodec) Option { _ = "STUB: not implemented"; return *new(Option) }

func defaultTimeoutOption() Option { _ = "STUB: not implemented"; return *new(Option) }
