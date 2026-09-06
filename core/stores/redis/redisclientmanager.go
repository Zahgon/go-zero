package redis

import (
	"runtime"

	red "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/syncx"
)

const (
	defaultDatabase = 0
	maxRetries      = 3
	idleConns       = 8
)

var (
	clientManager = syncx.NewResourceManager()

	nodePoolSize = 10 * runtime.GOMAXPROCS(0)
)

func getClient(r *Redis) (*red.Client, error) { _ = "STUB: not implemented"; return nil, nil }
