package cache

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/mathx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
)

const (
	notFoundPlaceholder = "*"

	expiryDeviation = 0.05
)

var errPlaceholder = errors.New("placeholder")

type cacheNode struct {
	rds            *redis.Redis
	expiry         time.Duration
	notFoundExpiry time.Duration
	barrier        syncx.SingleFlight
	r              *rand.Rand
	lock           *sync.Mutex
	unstableExpiry mathx.Unstable
	stat           *Stat
	errNotFound    error
}

func NewNode(rds *redis.Redis, barrier syncx.SingleFlight, st *Stat,
	errNotFound error, opts ...Option) Cache {
	_ = "STUB: not implemented"
	return *new(Cache)
}

func (c cacheNode) Del(keys ...string) error { _ = "STUB: not implemented"; return nil }

func (c cacheNode) DelCtx(ctx context.Context, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) Get(key string, val any) error { _ = "STUB: not implemented"; return nil }

func (c cacheNode) GetCtx(ctx context.Context, key string, val any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) IsNotFound(err error) bool { _ = "STUB: not implemented"; return false }

func (c cacheNode) Set(key string, val any) error { _ = "STUB: not implemented"; return nil }

func (c cacheNode) SetCtx(ctx context.Context, key string, val any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) SetWithExpire(key string, val any, expire time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) SetWithExpireCtx(ctx context.Context, key string, val any,
	expire time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) String() string { _ = "STUB: not implemented"; return "" }

func (c cacheNode) Take(val any, key string, query func(val any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) TakeCtx(ctx context.Context, val any, key string,
	query func(val any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) TakeWithExpire(val any, key string, query func(val any,
	expire time.Duration) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) TakeWithExpireCtx(ctx context.Context, val any, key string,
	query func(val any, expire time.Duration) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) aroundDuration(duration time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c cacheNode) asyncRetryDelCache(keys ...string) { _ = "STUB: not implemented"; return }

func (c cacheNode) doGetCache(ctx context.Context, key string, v any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) doTake(ctx context.Context, v any, key string,
	query func(v any) error, cacheVal func(v any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) processCache(ctx context.Context, key, data string, v any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c cacheNode) setCacheWithNotFound(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}
