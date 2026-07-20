package zrpc

import (
	"time"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc/internal"
)

type (
	ClientMiddlewaresConf = internal.ClientMiddlewaresConf

	ServerMiddlewaresConf = internal.ServerMiddlewaresConf

	StatConf = internal.StatConf

	MethodTimeoutConf = internal.MethodTimeoutConf

	RpcClientConf struct {
		Etcd          discov.EtcdConf `json:",optional,inherit"`
		Endpoints     []string        `json:",optional"`
		Target        string          `json:",optional"`
		App           string          `json:",optional"`
		Token         string          `json:",optional"`
		NonBlock      bool            `json:",default=true"`
		Timeout       int64           `json:",default=2000"`
		KeepaliveTime time.Duration   `json:",optional"`
		Middlewares   ClientMiddlewaresConf
		BalancerName  string `json:",default=p2c_ewma"`
	}

	RpcServerConf struct {
		service.ServiceConf
		ListenOn      string
		Etcd          discov.EtcdConf    `json:",optional,inherit"`
		Auth          bool               `json:",optional"`
		Redis         redis.RedisKeyConf `json:",optional"`
		StrictControl bool               `json:",optional"`

		Timeout      int64 `json:",default=2000"`
		CpuThreshold int64 `json:",default=900,range=[0:1000)"`

		Health      bool `json:",default=true"`
		Middlewares ServerMiddlewaresConf

		MethodTimeouts []MethodTimeoutConf `json:",optional"`
	}
)

func NewDirectClientConf(endpoints []string, app, token string) RpcClientConf {
	_ = "STUB: not implemented"
	return *new(RpcClientConf)
}

func NewEtcdClientConf(hosts []string, key, app, token string) RpcClientConf {
	_ = "STUB: not implemented"
	return *new(RpcClientConf)
}

func (sc RpcServerConf) HasEtcd() bool { _ = "STUB: not implemented"; return false }

func (sc RpcServerConf) Validate() error { _ = "STUB: not implemented"; return nil }

func (cc RpcClientConf) BuildTarget() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (cc RpcClientConf) HasCredential() bool { _ = "STUB: not implemented"; return false }
