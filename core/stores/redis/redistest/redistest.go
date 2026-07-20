package redistest

import (
	"testing"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

func CreateRedis(t *testing.T) *redis.Redis { _ = "STUB: not implemented"; return nil }

func CreateRedisWithClean(t *testing.T) (r *redis.Redis, clean func()) {
	_ = "STUB: not implemented"
	return nil, nil
}
