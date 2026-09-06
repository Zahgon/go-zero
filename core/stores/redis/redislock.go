package redis

import (
	"context"
	_ "embed"
	"math/rand"
	"time"
)

const (
	randomLen       = 16
	tolerance       = 500
	millisPerSecond = 1000
)

var (
	//go:embed lockscript.lua
	lockLuaScript string
	lockScript    = NewScript(lockLuaScript)

	//go:embed delscript.lua
	delLuaScript string
	delScript    = NewScript(delLuaScript)
)

type RedisLock struct {
	store   *Redis
	seconds uint32
	key     string
	id      string
}

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func NewRedisLock(store *Redis, key string) *RedisLock { _ = "STUB: not implemented"; return nil }

func (rl *RedisLock) Acquire() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (rl *RedisLock) AcquireCtx(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rl *RedisLock) Release() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (rl *RedisLock) ReleaseCtx(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rl *RedisLock) SetExpire(seconds int) { _ = "STUB: not implemented"; return }
