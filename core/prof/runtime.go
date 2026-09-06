package prof

import (
	"io"
	"time"
)

const (
	defaultInterval = time.Second * 5
	mega            = 1024 * 1024
)

func DisplayStats(interval ...time.Duration) { _ = "STUB: not implemented"; return }

func displayStatsWithWriter(writer io.Writer, interval ...time.Duration) {
	_ = "STUB: not implemented"
	return
}
