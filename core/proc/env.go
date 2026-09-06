package proc

import (
	"sync"
)

var (
	envs    = make(map[string]string)
	envLock sync.RWMutex
)

func Env(name string) string { _ = "STUB: not implemented"; return "" }

func EnvInt(name string) (int, bool) { _ = "STUB: not implemented"; return 0, false }
