package logx

import (
	"time"

	"github.com/zeromicro/go-zero/core/syncx"
)

type limitedExecutor struct {
	threshold time.Duration
	lastTime  *syncx.AtomicDuration
	discarded uint32
}

func newLimitedExecutor(milliseconds int) *limitedExecutor { _ = "STUB: not implemented"; return nil }

func (le *limitedExecutor) logOrDiscard(execute func()) { _ = "STUB: not implemented"; return }
