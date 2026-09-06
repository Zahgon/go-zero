package prometheus

import (
	"sync"

	"github.com/zeromicro/go-zero/core/syncx"
)

var (
	once    sync.Once
	enabled syncx.AtomicBool
)

func Enabled() bool { _ = "STUB: not implemented"; return false }

func Enable() { _ = "STUB: not implemented"; return }

func StartAgent(c Config) { _ = "STUB: not implemented"; return }
