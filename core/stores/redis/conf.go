package redis

import (
	"errors"
	"time"
)

var (
	ErrEmptyHost = errors.New("empty redis host")

	ErrEmptyType = errors.New("empty redis type")

	ErrEmptyKey = errors.New("empty redis key")
)

type (
	RedisConf struct {
		Host     string
		Type     string `json:",default=node,options=node|cluster"`
		User     string `json:",optional"`
		Pass     string `json:",optional"`
		Tls      bool   `json:",optional"`
		NonBlock bool   `json:",default=true"`

		DisableIdentity bool `json:",default=false"`

		Protocol int `json:",default=3"`

		MaintNotifications string `json:",default=disabled,options=disabled|enabled|auto"`

		PingTimeout time.Duration `json:",default=1s"`
	}

	RedisKeyConf struct {
		RedisConf
		Key string
	}
)

func (rc RedisConf) NewRedis() *Redis { _ = "STUB: not implemented"; return nil }

func (rc RedisConf) Validate() error { _ = "STUB: not implemented"; return nil }

func (rkc RedisKeyConf) Validate() error { _ = "STUB: not implemented"; return nil }
