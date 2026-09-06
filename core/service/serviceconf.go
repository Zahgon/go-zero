package service

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/core/prometheus"
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/internal/devserver"
	"github.com/zeromicro/go-zero/internal/profiling"
)

const (
	DevMode = "dev"

	TestMode = "test"

	RtMode = "rt"

	PreMode = "pre"

	ProMode = "pro"
)

type (
	DevServerConfig = devserver.Config

	ServiceConf struct {
		Name       string
		Log        logx.LogConf
		Mode       string `json:",default=pro,options=dev|test|rt|pre|pro"`
		MetricsUrl string `json:",optional"`

		Prometheus prometheus.Config `json:",optional"`
		Telemetry  trace.Config      `json:",optional"`
		DevServer  DevServerConfig   `json:",optional"`
		Shutdown   proc.ShutdownConf `json:",optional"`

		Profiling profiling.Config `json:",optional"`
	}
)

func (sc ServiceConf) MustSetUp() { _ = "STUB: not implemented"; return }

func (sc ServiceConf) SetUp() error { _ = "STUB: not implemented"; return nil }

func (sc ServiceConf) initMode() { _ = "STUB: not implemented"; return }
