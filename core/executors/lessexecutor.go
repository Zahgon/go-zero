package executors

import (
	"time"

	"github.com/zeromicro/go-zero/core/syncx"
)

type LessExecutor struct {
	threshold time.Duration
	lastTime  *syncx.AtomicDuration
}

func NewLessExecutor(threshold time.Duration) *LessExecutor { _ = "STUB: not implemented"; return nil }

func (le *LessExecutor) DoOrDiscard(execute func()) bool { _ = "STUB: not implemented"; return false }
