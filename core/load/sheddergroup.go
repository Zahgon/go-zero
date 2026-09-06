package load

import (
	"github.com/zeromicro/go-zero/core/syncx"
)

type ShedderGroup struct {
	options []ShedderOption
	manager *syncx.ResourceManager
}

func NewShedderGroup(opts ...ShedderOption) *ShedderGroup { _ = "STUB: not implemented"; return nil }

func (g *ShedderGroup) GetShedder(key string) Shedder {
	_ = "STUB: not implemented"
	return *new(Shedder)
}

type nopCloser struct {
	Shedder
}

func (c nopCloser) Close() error { _ = "STUB: not implemented"; return nil }
