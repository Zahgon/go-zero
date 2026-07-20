package gateway

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type (
	GatewayConf struct {
		rest.RestConf
		Upstreams []Upstream
	}

	HttpClientConf struct {
		Target  string
		Prefix  string `json:",optional"`
		Timeout int64  `json:",default=3000"`
	}

	RouteMapping struct {
		Method string

		Path string

		RpcPath string `json:",optional"`
	}

	Upstream struct {
		Name string `json:",optional"`

		Grpc *zrpc.RpcClientConf `json:",optional"`

		Http *HttpClientConf `json:",optional=!grpc"`

		ProtoSets []string `json:",optional"`

		Mappings []RouteMapping `json:",optional"`
	}
)
