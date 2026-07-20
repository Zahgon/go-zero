package subscriber

import (
	"sync"
	"sync/atomic"

	"github.com/zeromicro/go-zero/core/discov"
)

type (
	etcdSubscriber struct {
		*discov.Subscriber
	}

	EtcdConf = discov.EtcdConf
)

func MustNewEtcdSubscriber(conf EtcdConf) Subscriber {
	_ = "STUB: not implemented"
	return *new(Subscriber)
}

func NewEtcdSubscriber(conf EtcdConf) (Subscriber, error) {
	_ = "STUB: not implemented"
	return *new(Subscriber), nil
}

func buildSubOptions(conf EtcdConf) []discov.SubOption { _ = "STUB: not implemented"; return nil }

func (s *etcdSubscriber) AddListener(listener func()) error { _ = "STUB: not implemented"; return nil }

func (s *etcdSubscriber) Value() (string, error) { _ = "STUB: not implemented"; return "", nil }

type container struct {
	value     atomic.Value
	listeners []func()
	lock      sync.Mutex
}

func newContainer() *container { _ = "STUB: not implemented"; return nil }

func (c *container) OnAdd(kv discov.KV) { _ = "STUB: not implemented"; return }

func (c *container) OnDelete(_ discov.KV) { _ = "STUB: not implemented"; return }

func (c *container) AddListener(listener func()) { _ = "STUB: not implemented"; return }

func (c *container) GetValues() []string { _ = "STUB: not implemented"; return nil }

func (c *container) notifyChange() { _ = "STUB: not implemented"; return }
