package cache

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/syncx"
)

type (
	Cache interface {
		Del(keys ...string) error

		DelCtx(ctx context.Context, keys ...string) error

		Get(key string, val any) error

		GetCtx(ctx context.Context, key string, val any) error

		IsNotFound(err error) bool

		Set(key string, val any) error

		SetCtx(ctx context.Context, key string, val any) error

		SetWithExpire(key string, val any, expire time.Duration) error

		SetWithExpireCtx(ctx context.Context, key string, val any, expire time.Duration) error

		Take(val any, key string, query func(val any) error) error

		TakeCtx(ctx context.Context, val any, key string, query func(val any) error) error

		TakeWithExpire(val any, key string, query func(val any, expire time.Duration) error) error

		TakeWithExpireCtx(ctx context.Context, val any, key string,
			query func(val any, expire time.Duration) error) error
	}

	cacheCluster struct {
		dispatcher  *hash.ConsistentHash
		errNotFound error
	}
)

func New(c ClusterConf, barrier syncx.SingleFlight, st *Stat, errNotFound error,
	opts ...Option) Cache {
	_ = "STUB: not implemented"
	return *new(Cache)
}

func (cc cacheCluster) Del(keys ...string) error { _ = "STUB: not implemented"; return nil }

func (cc cacheCluster) DelCtx(ctx context.Context, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc cacheCluster) Get(key string, val any) error { _ = "STUB: not implemented"; return nil }

func (cc cacheCluster) GetCtx(ctx context.Context, key string, val any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc cacheCluster) IsNotFound(err error) bool { _ = "STUB: not implemented"; return false }

func (cc cacheCluster) Set(key string, val any) error { _ = "STUB: not implemented"; return nil }

func (cc cacheCluster) SetCtx(ctx context.Context, key string, val any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc cacheCluster) SetWithExpire(key string, val any, expire time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc cacheCluster) SetWithExpireCtx(ctx context.Context, key string, val any, expire time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc cacheCluster) Take(val any, key string, query func(val any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc cacheCluster) TakeCtx(ctx context.Context, val any, key string, query func(val any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc cacheCluster) TakeWithExpire(val any, key string, query func(val any, expire time.Duration) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc cacheCluster) TakeWithExpireCtx(ctx context.Context, val any, key string, query func(val any, expire time.Duration) error) error {
	_ = "STUB: not implemented"
	return nil
}
