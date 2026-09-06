//go:build linux

package stat

import (
	"flag"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/executors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
)

const (
	clusterNameKey = "CLUSTER_NAME"
	testEnv        = "test.v"
)

var (
	reporter     = logx.Alert
	lock         sync.RWMutex
	lessExecutor = executors.NewLessExecutor(time.Minute * 5)
	dropped      int32
	clusterName  = proc.Env(clusterNameKey)
)

func init() {
	if flag.Lookup(testEnv) != nil {
		SetReporter(nil)
	}
}

func Report(msg string) { _ = "STUB: not implemented"; return }

func SetReporter(fn func(string)) { _ = "STUB: not implemented"; return }
