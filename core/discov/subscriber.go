package discov

import (
	"sync"
	"sync/atomic"

	"github.com/zeromicro/go-zero/core/discov/internal"
	"github.com/zeromicro/go-zero/core/syncx"
)

type (
	SubOption func(sub *Subscriber)

	Subscriber struct {
		endpoints  []string
		exclusive  bool
		key        string
		exactMatch bool
		items      Container
	}
	KV = internal.KV
)

func NewSubscriber(endpoints []string, key string, opts ...SubOption) (*Subscriber, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Subscriber) AddListener(listener func()) { _ = "STUB: not implemented"; return }

func (s *Subscriber) Close() { _ = "STUB: not implemented"; return }

func (s *Subscriber) Values() []string { _ = "STUB: not implemented"; return nil }

func Exclusive() SubOption { _ = "STUB: not implemented"; return *new(SubOption) }

func WithExactMatch() SubOption { _ = "STUB: not implemented"; return *new(SubOption) }

func WithSubEtcdAccount(user, pass string) SubOption {
	_ = "STUB: not implemented"
	return *new(SubOption)
}

func WithSubEtcdTLS(certFile, certKeyFile, caFile string, insecureSkipVerify bool) SubOption {
	_ = "STUB: not implemented"
	return *new(SubOption)
}

func WithContainer(container Container) SubOption {
	_ = "STUB: not implemented"
	return *new(SubOption)
}

type (
	Container interface {
		OnAdd(kv internal.KV)
		OnDelete(kv internal.KV)
		AddListener(listener func())
		GetValues() []string
	}

	container struct {
		exclusive bool
		values    map[string][]string
		mapping   map[string]string
		snapshot  atomic.Value
		dirty     *syncx.AtomicBool
		listeners []func()
		lock      sync.Mutex
	}
)

func newContainer(exclusive bool) *container { _ = "STUB: not implemented"; return nil }

func (c *container) OnAdd(kv internal.KV) { _ = "STUB: not implemented"; return }

func (c *container) OnDelete(kv internal.KV) { _ = "STUB: not implemented"; return }

func (c *container) addKv(key, value string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *container) AddListener(listener func()) { _ = "STUB: not implemented"; return }

func (c *container) doRemoveKey(key string) { _ = "STUB: not implemented"; return }

func (c *container) GetValues() []string { _ = "STUB: not implemented"; return nil }

func (c *container) notifyChange() { _ = "STUB: not implemented"; return }

func (c *container) removeKey(key string) { _ = "STUB: not implemented"; return }
