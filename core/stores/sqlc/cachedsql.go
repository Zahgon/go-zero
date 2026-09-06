package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
)

const cacheSafeGapBetweenIndexAndPrimary = time.Second * 5

var (
	ErrNotFound = sqlx.ErrNotFound

	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("sqlc")
)

type (
	ExecFn func(conn sqlx.SqlConn) (sql.Result, error)

	ExecCtxFn func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error)

	IndexQueryFn func(conn sqlx.SqlConn, v any) (any, error)

	IndexQueryCtxFn func(ctx context.Context, conn sqlx.SqlConn, v any) (any, error)

	PrimaryQueryFn func(conn sqlx.SqlConn, v, primary any) error

	PrimaryQueryCtxFn func(ctx context.Context, conn sqlx.SqlConn, v, primary any) error

	QueryFn func(conn sqlx.SqlConn, v any) error

	QueryCtxFn func(ctx context.Context, conn sqlx.SqlConn, v any) error

	CachedConn struct {
		db    sqlx.SqlConn
		cache cache.Cache
	}
)

func NewConn(db sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CachedConn {
	_ = "STUB: not implemented"
	return *new(CachedConn)
}

func NewConnWithCache(db sqlx.SqlConn, c cache.Cache) CachedConn {
	_ = "STUB: not implemented"
	return *new(CachedConn)
}

func NewNodeConn(db sqlx.SqlConn, rds *redis.Redis, opts ...cache.Option) CachedConn {
	_ = "STUB: not implemented"
	return *new(CachedConn)
}

func (cc CachedConn) DelCache(keys ...string) error { _ = "STUB: not implemented"; return nil }

func (cc CachedConn) DelCacheCtx(ctx context.Context, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) GetCache(key string, v any) error { _ = "STUB: not implemented"; return nil }

func (cc CachedConn) GetCacheCtx(ctx context.Context, key string, v any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) Exec(exec ExecFn, keys ...string) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (cc CachedConn) ExecCtx(ctx context.Context, exec ExecCtxFn, keys ...string) (
	sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (cc CachedConn) ExecNoCache(q string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (cc CachedConn) ExecNoCacheCtx(ctx context.Context, q string, args ...any) (
	sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (cc CachedConn) QueryRow(v any, key string, query QueryFn) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) QueryRowCtx(ctx context.Context, v any, key string, query QueryCtxFn) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) QueryRowIndex(v any, key string, keyer func(primary any) string,
	indexQuery IndexQueryFn, primaryQuery PrimaryQueryFn) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) QueryRowIndexCtx(ctx context.Context, v any, key string,
	keyer func(primary any) string, indexQuery IndexQueryCtxFn,
	primaryQuery PrimaryQueryCtxFn) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) QueryRowNoCache(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) QueryRowNoCacheCtx(ctx context.Context, v any, q string,
	args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) QueryRowPartialNoCache(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) QueryRowPartialNoCacheCtx(ctx context.Context, v any, q string,
	args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) QueryRowsNoCache(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) QueryRowsNoCacheCtx(ctx context.Context, v any, q string,
	args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) QueryRowsPartialNoCache(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) QueryRowsPartialNoCacheCtx(ctx context.Context, v any, q string,
	args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) SetCache(key string, val any) error { _ = "STUB: not implemented"; return nil }

func (cc CachedConn) SetCacheCtx(ctx context.Context, key string, val any) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) SetCacheWithExpire(key string, val any, expire time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) SetCacheWithExpireCtx(ctx context.Context, key string, val any,
	expire time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) Transact(fn func(sqlx.Session) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) TransactCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc CachedConn) WithSession(session sqlx.Session) CachedConn {
	_ = "STUB: not implemented"
	return *new(CachedConn)
}
