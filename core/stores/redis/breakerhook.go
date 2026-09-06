package redis

import (
	red "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/breaker"
	"github.com/zeromicro/go-zero/core/lang"
)

var ignoreCmds = map[string]lang.PlaceholderType{
	"blpop": {},
	"hello": {},
}

type breakerHook struct {
	brk breaker.Breaker
}

func (h breakerHook) DialHook(next red.DialHook) red.DialHook {
	_ = "STUB: not implemented"
	return *new(red.DialHook)
}

func (h breakerHook) ProcessHook(next red.ProcessHook) red.ProcessHook {
	_ = "STUB: not implemented"
	return *new(red.ProcessHook)
}

func (h breakerHook) ProcessPipelineHook(next red.ProcessPipelineHook) red.ProcessPipelineHook {
	_ = "STUB: not implemented"
	return *new(red.ProcessPipelineHook)
}
