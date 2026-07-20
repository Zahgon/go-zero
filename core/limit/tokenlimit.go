package limit

import (
	"context"
	_ "embed"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
	xrate "golang.org/x/time/rate"
)

const (
	tokenFormat     = "{%s}.tokens"
	timestampFormat = "{%s}.ts"
	pingInterval    = time.Millisecond * 100
)

var (
	//go:embed tokenscript.lua
	tokenLuaScript string
	tokenScript    = redis.NewScript(tokenLuaScript)
)

type TokenLimiter struct {
	rate           int
	burst          int
	store          *redis.Redis
	tokenKey       string
	timestampKey   string
	rescueLock     sync.Mutex
	redisAlive     uint32
	monitorStarted bool
	rescueLimiter  *xrate.Limiter
}

func NewTokenLimiter(rate, burst int, store *redis.Redis, key string) *TokenLimiter {
	_ = "STUB: not implemented"
	return nil
}

func (lim *TokenLimiter) Allow() bool { _ = "STUB: not implemented"; return false }

func (lim *TokenLimiter) AllowCtx(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (lim *TokenLimiter) AllowN(now time.Time, n int) bool { _ = "STUB: not implemented"; return false }

func (lim *TokenLimiter) AllowNCtx(ctx context.Context, now time.Time, n int) bool {
	_ = "STUB: not implemented"
	return false
}

func (lim *TokenLimiter) reserveN(ctx context.Context, now time.Time, n int) bool {
	_ = "STUB: not implemented"
	return false
}

func (lim *TokenLimiter) startMonitor() { _ = "STUB: not implemented"; return }

func (lim *TokenLimiter) waitForRedis() { _ = "STUB: not implemented"; return }
