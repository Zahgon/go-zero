package redis

import (
	"sync"
	"sync/atomic"
)

var (
	once     sync.Once
	lock     sync.Mutex
	instance *ScriptCache
)

type (
	Map map[string]string

	ScriptCache struct {
		atomic.Value
	}
)

func GetScriptCache() *ScriptCache { _ = "STUB: not implemented"; return nil }

func (sc *ScriptCache) GetSha(script string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (sc *ScriptCache) SetSha(script, sha string) { _ = "STUB: not implemented"; return }
