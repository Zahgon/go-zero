package kv

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var ErrNoRedisNode = errors.New("no redis node")

type (
	Store interface {
		Decr(key string) (int64, error)
		DecrCtx(ctx context.Context, key string) (int64, error)
		Decrby(key string, decrement int64) (int64, error)
		DecrbyCtx(ctx context.Context, key string, decrement int64) (int64, error)
		Del(keys ...string) (int, error)
		DelCtx(ctx context.Context, keys ...string) (int, error)
		Eval(script, key string, args ...any) (any, error)
		EvalCtx(ctx context.Context, script, key string, args ...any) (any, error)
		Exists(key string) (bool, error)
		ExistsCtx(ctx context.Context, key string) (bool, error)
		Expire(key string, seconds int) error
		ExpireCtx(ctx context.Context, key string, seconds int) error
		Expireat(key string, expireTime int64) error
		ExpireatCtx(ctx context.Context, key string, expireTime int64) error
		Get(key string) (string, error)
		GetCtx(ctx context.Context, key string) (string, error)
		GetSet(key, value string) (string, error)
		GetSetCtx(ctx context.Context, key, value string) (string, error)
		Hdel(key, field string) (bool, error)
		HdelCtx(ctx context.Context, key, field string) (bool, error)
		Hexists(key, field string) (bool, error)
		HexistsCtx(ctx context.Context, key, field string) (bool, error)
		Hget(key, field string) (string, error)
		HgetCtx(ctx context.Context, key, field string) (string, error)
		Hgetall(key string) (map[string]string, error)
		HgetallCtx(ctx context.Context, key string) (map[string]string, error)
		Hincrby(key, field string, increment int) (int, error)
		HincrbyCtx(ctx context.Context, key, field string, increment int) (int, error)
		Hkeys(key string) ([]string, error)
		HkeysCtx(ctx context.Context, key string) ([]string, error)
		Hlen(key string) (int, error)
		HlenCtx(ctx context.Context, key string) (int, error)
		Hmget(key string, fields ...string) ([]string, error)
		HmgetCtx(ctx context.Context, key string, fields ...string) ([]string, error)
		Hset(key, field, value string) error
		HsetCtx(ctx context.Context, key, field, value string) error
		Hsetnx(key, field, value string) (bool, error)
		HsetnxCtx(ctx context.Context, key, field, value string) (bool, error)
		Hmset(key string, fieldsAndValues map[string]string) error
		HmsetCtx(ctx context.Context, key string, fieldsAndValues map[string]string) error
		Hvals(key string) ([]string, error)
		HvalsCtx(ctx context.Context, key string) ([]string, error)
		Incr(key string) (int64, error)
		IncrCtx(ctx context.Context, key string) (int64, error)
		Incrby(key string, increment int64) (int64, error)
		IncrbyCtx(ctx context.Context, key string, increment int64) (int64, error)
		Lindex(key string, index int64) (string, error)
		LindexCtx(ctx context.Context, key string, index int64) (string, error)
		Llen(key string) (int, error)
		LlenCtx(ctx context.Context, key string) (int, error)
		Lpop(key string) (string, error)
		LpopCtx(ctx context.Context, key string) (string, error)
		Lpush(key string, values ...any) (int, error)
		LpushCtx(ctx context.Context, key string, values ...any) (int, error)
		Lrange(key string, start, stop int) ([]string, error)
		LrangeCtx(ctx context.Context, key string, start, stop int) ([]string, error)
		Lrem(key string, count int, value string) (int, error)
		LremCtx(ctx context.Context, key string, count int, value string) (int, error)
		Persist(key string) (bool, error)
		PersistCtx(ctx context.Context, key string) (bool, error)
		Pfadd(key string, values ...any) (bool, error)
		PfaddCtx(ctx context.Context, key string, values ...any) (bool, error)
		Pfcount(key string) (int64, error)
		PfcountCtx(ctx context.Context, key string) (int64, error)
		Rpush(key string, values ...any) (int, error)
		RpushCtx(ctx context.Context, key string, values ...any) (int, error)
		Sadd(key string, values ...any) (int, error)
		SaddCtx(ctx context.Context, key string, values ...any) (int, error)
		Scard(key string) (int64, error)
		ScardCtx(ctx context.Context, key string) (int64, error)
		Set(key, value string) error
		SetCtx(ctx context.Context, key, value string) error
		Setex(key, value string, seconds int) error
		SetexCtx(ctx context.Context, key, value string, seconds int) error
		Setnx(key, value string) (bool, error)
		SetnxCtx(ctx context.Context, key, value string) (bool, error)
		SetnxEx(key, value string, seconds int) (bool, error)
		SetnxExCtx(ctx context.Context, key, value string, seconds int) (bool, error)
		Sismember(key string, value any) (bool, error)
		SismemberCtx(ctx context.Context, key string, value any) (bool, error)
		Smembers(key string) ([]string, error)
		SmembersCtx(ctx context.Context, key string) ([]string, error)
		Spop(key string) (string, error)
		SpopCtx(ctx context.Context, key string) (string, error)
		Srandmember(key string, count int) ([]string, error)
		SrandmemberCtx(ctx context.Context, key string, count int) ([]string, error)
		Srem(key string, values ...any) (int, error)
		SremCtx(ctx context.Context, key string, values ...any) (int, error)
		Sscan(key string, cursor uint64, match string, count int64) (keys []string, cur uint64, err error)
		SscanCtx(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, cur uint64, err error)
		Ttl(key string) (int, error)
		TtlCtx(ctx context.Context, key string) (int, error)
		Zadd(key string, score int64, value string) (bool, error)
		ZaddFloat(key string, score float64, value string) (bool, error)
		ZaddCtx(ctx context.Context, key string, score int64, value string) (bool, error)
		ZaddFloatCtx(ctx context.Context, key string, score float64, value string) (bool, error)
		Zadds(key string, ps ...redis.Pair) (int64, error)
		ZaddsCtx(ctx context.Context, key string, ps ...redis.Pair) (int64, error)
		Zcard(key string) (int, error)
		ZcardCtx(ctx context.Context, key string) (int, error)
		Zcount(key string, start, stop int64) (int, error)
		ZcountCtx(ctx context.Context, key string, start, stop int64) (int, error)
		Zincrby(key string, increment int64, field string) (int64, error)
		ZincrbyCtx(ctx context.Context, key string, increment int64, field string) (int64, error)
		Zrange(key string, start, stop int64) ([]string, error)
		ZrangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error)
		ZrangeWithScores(key string, start, stop int64) ([]redis.Pair, error)
		ZrangeWithScoresCtx(ctx context.Context, key string, start, stop int64) ([]redis.Pair, error)
		ZrangebyscoreWithScores(key string, start, stop int64) ([]redis.Pair, error)
		ZrangebyscoreWithScoresCtx(ctx context.Context, key string, start, stop int64) ([]redis.Pair, error)
		ZrangebyscoreWithScoresAndLimit(key string, start, stop int64, page, size int) ([]redis.Pair, error)
		ZrangebyscoreWithScoresAndLimitCtx(ctx context.Context, key string, start, stop int64, page, size int) ([]redis.Pair, error)
		Zrank(key, field string) (int64, error)
		ZrankCtx(ctx context.Context, key, field string) (int64, error)
		Zrem(key string, values ...any) (int, error)
		ZremCtx(ctx context.Context, key string, values ...any) (int, error)
		Zremrangebyrank(key string, start, stop int64) (int, error)
		ZremrangebyrankCtx(ctx context.Context, key string, start, stop int64) (int, error)
		Zremrangebyscore(key string, start, stop int64) (int, error)
		ZremrangebyscoreCtx(ctx context.Context, key string, start, stop int64) (int, error)
		Zrevrange(key string, start, stop int64) ([]string, error)
		ZrevrangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error)
		ZrevrangebyscoreWithScores(key string, start, stop int64) ([]redis.Pair, error)
		ZrevrangebyscoreWithScoresCtx(ctx context.Context, key string, start, stop int64) ([]redis.Pair, error)
		ZrevrangebyscoreWithScoresAndLimit(key string, start, stop int64, page, size int) ([]redis.Pair, error)
		ZrevrangebyscoreWithScoresAndLimitCtx(ctx context.Context, key string, start, stop int64, page, size int) ([]redis.Pair, error)
		Zscore(key, value string) (int64, error)
		ZscoreCtx(ctx context.Context, key, value string) (int64, error)
		Zrevrank(key, field string) (int64, error)
		ZrevrankCtx(ctx context.Context, key, field string) (int64, error)
	}

	clusterStore struct {
		dispatcher *hash.ConsistentHash
	}
)

func NewStore(c KvConf) Store { _ = "STUB: not implemented"; return *new(Store) }

func (cs clusterStore) Decr(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cs clusterStore) DecrCtx(ctx context.Context, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Decrby(key string, decrement int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) DecrbyCtx(ctx context.Context, key string, decrement int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Del(keys ...string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cs clusterStore) DelCtx(ctx context.Context, keys ...string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Eval(script, key string, args ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (cs clusterStore) EvalCtx(ctx context.Context, script, key string, args ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (cs clusterStore) Exists(key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) ExistsCtx(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) Expire(key string, seconds int) error { _ = "STUB: not implemented"; return nil }

func (cs clusterStore) ExpireCtx(ctx context.Context, key string, seconds int) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs clusterStore) Expireat(key string, expireTime int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs clusterStore) ExpireatCtx(ctx context.Context, key string, expireTime int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs clusterStore) Get(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (cs clusterStore) GetCtx(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cs clusterStore) Hdel(key, field string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) HdelCtx(ctx context.Context, key, field string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) Hexists(key, field string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) HexistsCtx(ctx context.Context, key, field string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) Hget(key, field string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cs clusterStore) HgetCtx(ctx context.Context, key, field string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cs clusterStore) Hgetall(key string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) HgetallCtx(ctx context.Context, key string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) Hincrby(key, field string, increment int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) HincrbyCtx(ctx context.Context, key, field string, increment int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Hkeys(key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) HkeysCtx(ctx context.Context, key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) Hlen(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cs clusterStore) HlenCtx(ctx context.Context, key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Hmget(key string, fields ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) HmgetCtx(ctx context.Context, key string, fields ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) Hset(key, field, value string) error { _ = "STUB: not implemented"; return nil }

func (cs clusterStore) HsetCtx(ctx context.Context, key, field, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs clusterStore) Hsetnx(key, field, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) HsetnxCtx(ctx context.Context, key, field, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) Hmset(key string, fieldsAndValues map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs clusterStore) HmsetCtx(ctx context.Context, key string, fieldsAndValues map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs clusterStore) Hvals(key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) HvalsCtx(ctx context.Context, key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) Incr(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cs clusterStore) IncrCtx(ctx context.Context, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Incrby(key string, increment int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) IncrbyCtx(ctx context.Context, key string, increment int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Llen(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cs clusterStore) LlenCtx(ctx context.Context, key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Lindex(key string, index int64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cs clusterStore) LindexCtx(ctx context.Context, key string, index int64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cs clusterStore) Lpop(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (cs clusterStore) LpopCtx(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cs clusterStore) Lpush(key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) LpushCtx(ctx context.Context, key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Lrange(key string, start, stop int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) LrangeCtx(ctx context.Context, key string, start, stop int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) Lrem(key string, count int, value string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) LremCtx(ctx context.Context, key string, count int, value string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Persist(key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) PersistCtx(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) Pfadd(key string, values ...any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) PfaddCtx(ctx context.Context, key string, values ...any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) Pfcount(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cs clusterStore) PfcountCtx(ctx context.Context, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Rpush(key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) RpushCtx(ctx context.Context, key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Sadd(key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) SaddCtx(ctx context.Context, key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Scard(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cs clusterStore) ScardCtx(ctx context.Context, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Set(key, value string) error { _ = "STUB: not implemented"; return nil }

func (cs clusterStore) SetCtx(ctx context.Context, key, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs clusterStore) Setex(key, value string, seconds int) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs clusterStore) SetexCtx(ctx context.Context, key, value string, seconds int) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs clusterStore) Setnx(key, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) SetnxCtx(ctx context.Context, key, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) SetnxEx(key, value string, seconds int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) SetnxExCtx(ctx context.Context, key, value string, seconds int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) GetSet(key, value string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cs clusterStore) GetSetCtx(ctx context.Context, key, value string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cs clusterStore) Sismember(key string, value any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) SismemberCtx(ctx context.Context, key string, value any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) Smembers(key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) SmembersCtx(ctx context.Context, key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) Spop(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (cs clusterStore) SpopCtx(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cs clusterStore) Srandmember(key string, count int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) SrandmemberCtx(ctx context.Context, key string, count int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) Srem(key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) SremCtx(ctx context.Context, key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Sscan(key string, cursor uint64, match string, count int64) (
	keys []string, cur uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (cs clusterStore) SscanCtx(ctx context.Context, key string, cursor uint64, match string, count int64) (
	keys []string, cur uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (cs clusterStore) Ttl(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cs clusterStore) TtlCtx(ctx context.Context, key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Zadd(key string, score int64, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) ZaddFloat(key string, score float64, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) ZaddCtx(ctx context.Context, key string, score int64, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) ZaddFloatCtx(ctx context.Context, key string, score float64, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cs clusterStore) Zadds(key string, ps ...redis.Pair) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) ZaddsCtx(ctx context.Context, key string, ps ...redis.Pair) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Zcard(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cs clusterStore) ZcardCtx(ctx context.Context, key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Zcount(key string, start, stop int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) ZcountCtx(ctx context.Context, key string, start, stop int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Zincrby(key string, increment int64, field string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) ZincrbyCtx(ctx context.Context, key string, increment int64, field string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Zrank(key, field string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) ZrankCtx(ctx context.Context, key, field string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Zrange(key string, start, stop int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrangeWithScores(key string, start, stop int64) ([]redis.Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrangeWithScoresCtx(ctx context.Context, key string, start, stop int64) ([]redis.Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrangebyscoreWithScores(key string, start, stop int64) ([]redis.Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrangebyscoreWithScoresCtx(ctx context.Context, key string, start, stop int64) ([]redis.Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrangebyscoreWithScoresAndLimit(key string, start, stop int64, page, size int) (
	[]redis.Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrangebyscoreWithScoresAndLimitCtx(ctx context.Context, key string, start, stop int64, page, size int) (
	[]redis.Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) Zrem(key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) ZremCtx(ctx context.Context, key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Zremrangebyrank(key string, start, stop int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) ZremrangebyrankCtx(ctx context.Context, key string, start, stop int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Zremrangebyscore(key string, start, stop int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) ZremrangebyscoreCtx(ctx context.Context, key string, start, stop int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Zrevrange(key string, start, stop int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrevrangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrevrangebyscoreWithScores(key string, start, stop int64) ([]redis.Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrevrangebyscoreWithScoresCtx(ctx context.Context, key string, start, stop int64) ([]redis.Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrevrangebyscoreWithScoresAndLimit(key string, start, stop int64, page, size int) (
	[]redis.Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) ZrevrangebyscoreWithScoresAndLimitCtx(ctx context.Context, key string, start, stop int64, page, size int) (
	[]redis.Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs clusterStore) Zrevrank(key, field string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) ZrevrankCtx(ctx context.Context, key, field string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) Zscore(key, value string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) ZscoreCtx(ctx context.Context, key, value string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs clusterStore) getRedis(key string) (*redis.Redis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
