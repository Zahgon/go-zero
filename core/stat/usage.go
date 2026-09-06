package stat

import (
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/stat/internal"
	"github.com/zeromicro/go-zero/core/threading"
)

const (
	cpuRefreshInterval = time.Millisecond * 250
	allRefreshInterval = time.Minute

	beta = 0.95
)

var cpuUsage int64

func init() {
	go func() {
		cpuTicker := time.NewTicker(cpuRefreshInterval)
		defer cpuTicker.Stop()
		allTicker := time.NewTicker(allRefreshInterval)
		defer allTicker.Stop()

		for {
			select {
			case <-cpuTicker.C:
				threading.RunSafe(func() {
					curUsage := internal.RefreshCpu()
					prevUsage := atomic.LoadInt64(&cpuUsage)

					usage := int64(float64(prevUsage)*beta + float64(curUsage)*(1-beta))
					atomic.StoreInt64(&cpuUsage, usage)
				})
			case <-allTicker.C:
				if logEnabled.True() {
					printUsage()
				}
			}
		}
	}()
}

func CpuUsage() int64 { _ = "STUB: not implemented"; return 0 }

func bToMb(b uint64) float32 { _ = "STUB: not implemented"; return 0 }

func printUsage() { _ = "STUB: not implemented"; return }
