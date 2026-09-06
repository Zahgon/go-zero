package rest

import (
	"time"

	"github.com/zeromicro/go-zero/core/service"
)

type (
	MiddlewaresConf struct {
		Trace      bool `json:",default=true"`
		Log        bool `json:",default=true"`
		Prometheus bool `json:",default=true"`
		MaxConns   bool `json:",default=true"`
		Breaker    bool `json:",default=true"`
		Shedding   bool `json:",default=true"`
		Timeout    bool `json:",default=true"`
		Recover    bool `json:",default=true"`
		Metrics    bool `json:",default=true"`
		MaxBytes   bool `json:",default=true"`
		Gunzip     bool `json:",default=true"`
	}

	PrivateKeyConf struct {
		Fingerprint string
		KeyFile     string
	}

	SignatureConf struct {
		Strict      bool          `json:",default=false"`
		Expiry      time.Duration `json:",default=1h"`
		PrivateKeys []PrivateKeyConf
	}

	RestConf struct {
		service.ServiceConf
		Host     string `json:",default=0.0.0.0"`
		Port     int
		CertFile string `json:",optional"`
		KeyFile  string `json:",optional"`
		Verbose  bool   `json:",optional"`
		MaxConns int    `json:",default=10000"`
		MaxBytes int64  `json:",default=1048576"`

		Timeout      int64         `json:",default=3000"`
		CpuThreshold int64         `json:",default=900,range=[0:1000)"`
		Signature    SignatureConf `json:",optional"`

		Middlewares MiddlewaresConf

		TraceIgnorePaths []string `json:",optional"`
	}
)
