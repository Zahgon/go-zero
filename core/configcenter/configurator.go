package configurator

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
)

var (
	errEmptyConfig            = errors.New("empty config value")
	errMissingUnmarshalerType = errors.New("missing unmarshaler type")
)

type Configurator[T any] interface {
	GetConfig() (T, error)

	AddListener(listener func())
}

type (
	Config struct {
		Type string `json:",default=yaml,options=[yaml,json,toml]"`

		Log bool `json:",default=true"`
	}

	configCenter[T any] struct {
		conf        Config
		unmarshaler LoaderFn
		subscriber  subscriber.Subscriber
		listeners   []func()
		lock        sync.Mutex
		snapshot    atomic.Value
	}

	value[T any] struct {
		data        string
		marshalData T
		err         error
	}
)

var _ Configurator[any] = (*configCenter[any])(nil)

func MustNewConfigCenter[T any](c Config, subscriber subscriber.Subscriber) Configurator[T] {
	_ = "STUB: not implemented"
	return nil
}

func NewConfigCenter[T any](c Config, subscriber subscriber.Subscriber) (Configurator[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *configCenter[T]) AddListener(listener func()) { _ = "STUB: not implemented"; return }

func (c *configCenter[T]) GetConfig() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (c *configCenter[T]) Value() string { _ = "STUB: not implemented"; return "" }

func (c *configCenter[T]) loadConfig() error { _ = "STUB: not implemented"; return nil }

func (c *configCenter[T]) onChange() { _ = "STUB: not implemented"; return }

func (c *configCenter[T]) value() *value[T] { _ = "STUB: not implemented"; return nil }

func (c *configCenter[T]) genValue(data string) *value[T] { _ = "STUB: not implemented"; return nil }
