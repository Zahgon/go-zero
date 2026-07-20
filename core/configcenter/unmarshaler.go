package configurator

import (
	"sync"

	"github.com/zeromicro/go-zero/core/conf"
)

var registry = &unmarshalerRegistry{
	unmarshalers: map[string]LoaderFn{
		"json": conf.LoadFromJsonBytes,
		"toml": conf.LoadFromTomlBytes,
		"yaml": conf.LoadFromYamlBytes,
	},
}

type (
	LoaderFn func([]byte, any) error

	unmarshalerRegistry struct {
		unmarshalers map[string]LoaderFn
		mu           sync.RWMutex
	}
)

func RegisterUnmarshaler(name string, fn LoaderFn) { _ = "STUB: not implemented"; return }

func Unmarshaler(name string) (LoaderFn, bool) {
	_ = "STUB: not implemented"
	return *new(LoaderFn), false
}
