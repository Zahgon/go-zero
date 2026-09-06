package discov

import (
	"github.com/zeromicro/go-zero/core/discov/internal"
	"github.com/zeromicro/go-zero/core/lang"
	"github.com/zeromicro/go-zero/core/syncx"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type (
	PubOption func(client *Publisher)

	Publisher struct {
		endpoints  []string
		key        string
		fullKey    string
		id         int64
		value      string
		lease      clientv3.LeaseID
		quit       *syncx.DoneChan
		pauseChan  chan lang.PlaceholderType
		resumeChan chan lang.PlaceholderType
	}
)

func NewPublisher(endpoints []string, key, value string, opts ...PubOption) *Publisher {
	_ = "STUB: not implemented"
	return nil
}

func (p *Publisher) KeepAlive() error { _ = "STUB: not implemented"; return nil }

func (p *Publisher) Pause() { _ = "STUB: not implemented"; return }

func (p *Publisher) Resume() { _ = "STUB: not implemented"; return }

func (p *Publisher) Stop() { _ = "STUB: not implemented"; return }

func (p *Publisher) doKeepAlive() error { _ = "STUB: not implemented"; return nil }

func (p *Publisher) doRegister() (internal.EtcdClient, error) {
	_ = "STUB: not implemented"
	return *new(internal.EtcdClient), nil
}

func (p *Publisher) keepAliveAsync(cli internal.EtcdClient) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Publisher) register(client internal.EtcdClient) (clientv3.LeaseID, error) {
	_ = "STUB: not implemented"
	return *new(clientv3.LeaseID), nil
}

func (p *Publisher) revoke(cli internal.EtcdClient) { _ = "STUB: not implemented"; return }

func WithId(id int64) PubOption { _ = "STUB: not implemented"; return *new(PubOption) }

func WithPubEtcdAccount(user, pass string) PubOption {
	_ = "STUB: not implemented"
	return *new(PubOption)
}

func WithPubEtcdTLS(certFile, certKeyFile, caFile string, insecureSkipVerify bool) PubOption {
	_ = "STUB: not implemented"
	return *new(PubOption)
}
