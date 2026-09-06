package mcp

import (
	"time"

	"github.com/zeromicro/go-zero/rest"
)

type McpConf struct {
	rest.RestConf
	Mcp struct {
		Name string `json:",optional"`

		Version string `json:",default=1.0.0"`

		UseStreamable bool `json:",default=false"`

		SseEndpoint string `json:",default=/sse"`

		MessageEndpoint string `json:",default=/message"`

		Cors []string `json:",optional"`

		SseTimeout time.Duration `json:",default=24h"`

		MessageTimeout time.Duration `json:",default=30s"`
	}
}
