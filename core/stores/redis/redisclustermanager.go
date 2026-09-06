package redis

import (
	"runtime"

	red "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/syncx"
)

const addrSep = ","

var (
	clusterManager = syncx.NewResourceManager()

	clusterPoolSize = 5 * runtime.GOMAXPROCS(0)
)

func getCluster(r *Redis) (*red.ClusterClient, error) { _ = "STUB: not implemented"; return nil, nil }

func splitClusterAddrs(addr string) []string { _ = "STUB: not implemented"; return nil }
