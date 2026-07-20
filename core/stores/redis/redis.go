package redis

import (
	"context"
	"errors"
	"time"

	red "github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
	"github.com/zeromicro/go-zero/core/breaker"
	"github.com/zeromicro/go-zero/core/syncx"
)

const (
	ClusterType = "cluster"

	NodeType = "node"

	Nil = red.Nil

	blockingQueryTimeout = 5 * time.Second
	readWriteTimeout     = 2 * time.Second
	defaultSlowThreshold = time.Millisecond * 100
	defaultPingTimeout   = time.Second
)

var (
	ErrNilNode    = errors.New("nil redis node")
	slowThreshold = syncx.ForAtomicDuration(defaultSlowThreshold)
)

type (
	Option func(r *Redis)

	Pair struct {
		Key   string
		Score int64
	}

	FloatPair struct {
		Key   string
		Score float64
	}

	Redis struct {
		Addr               string
		Type               string
		User               string
		Pass               string
		protocol           int
		identity           bool
		maintNotifications maintnotifications.Mode
		tls                bool
		brk                breaker.Breaker
		hooks              []red.Hook
	}

	RedisNode interface {
		red.Cmdable
		Do(ctx context.Context, args ...any) *red.Cmd
	}

	GeoLocation = red.GeoLocation

	GeoRadiusQuery = red.GeoRadiusQuery

	GeoPos = red.GeoPos

	Pipeliner = red.Pipeliner

	Z = red.Z

	ZStore = red.ZStore

	IntCmd = red.IntCmd

	FloatCmd = red.FloatCmd

	StringCmd = red.StringCmd

	Script = red.Script

	Hook = red.Hook

	DialHook = red.DialHook

	ProcessHook = red.ProcessHook

	ProcessPipelineHook = red.ProcessPipelineHook

	Cmder = red.Cmder
)

func MustNewRedis(conf RedisConf, opts ...Option) *Redis { _ = "STUB: not implemented"; return nil }

func New(addr string, opts ...Option) *Redis { _ = "STUB: not implemented"; return nil }

func NewRedis(conf RedisConf, opts ...Option) (*Redis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewScript(script string) *Script { _ = "STUB: not implemented"; return nil }

func (s *Redis) BitCount(key string, start, end int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) BitCountCtx(ctx context.Context, key string, start, end int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) BitOpAnd(destKey string, keys ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) BitOpAndCtx(ctx context.Context, destKey string, keys ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) BitOpNot(destKey, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) BitOpNotCtx(ctx context.Context, destKey, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) BitOpOr(destKey string, keys ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) BitOpOrCtx(ctx context.Context, destKey string, keys ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) BitOpXor(destKey string, keys ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) BitOpXorCtx(ctx context.Context, destKey string, keys ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) BitPos(key string, bit, start, end int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) BitPosCtx(ctx context.Context, key string, bit, start, end int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Blpop(node RedisNode, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) BlpopCtx(ctx context.Context, node RedisNode, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) BlpopEx(node RedisNode, key string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func (s *Redis) BlpopExCtx(ctx context.Context, node RedisNode, key string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func (s *Redis) BlpopWithTimeout(node RedisNode, timeout time.Duration, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) BlpopWithTimeoutCtx(ctx context.Context, node RedisNode, timeout time.Duration,
	key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) Decr(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) DecrCtx(ctx context.Context, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Decrby(key string, decrement int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) DecrbyCtx(ctx context.Context, key string, decrement int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Del(keys ...string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) DelCtx(ctx context.Context, keys ...string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Do(args ...any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (s *Redis) DoCtx(ctx context.Context, args ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *Redis) Eval(script string, keys []string, args ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *Redis) EvalCtx(ctx context.Context, script string, keys []string,
	args ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *Redis) EvalSha(sha string, keys []string, args ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *Redis) EvalShaCtx(ctx context.Context, sha string, keys []string,
	args ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *Redis) Exists(key string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s *Redis) ExistsCtx(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) ExistsMany(keys ...string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) ExistsManyCtx(ctx context.Context, keys ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Expire(key string, seconds int) error { _ = "STUB: not implemented"; return nil }

func (s *Redis) ExpireCtx(ctx context.Context, key string, seconds int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Redis) Expireat(key string, expireTime int64) error { _ = "STUB: not implemented"; return nil }

func (s *Redis) ExpireatCtx(ctx context.Context, key string, expireTime int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Redis) GeoAdd(key string, geoLocation ...*GeoLocation) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) GeoAddCtx(ctx context.Context, key string, geoLocation ...*GeoLocation) (
	int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) GeoDist(key, member1, member2, unit string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) GeoDistCtx(ctx context.Context, key, member1, member2, unit string) (
	float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) GeoHash(key string, members ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) GeoHashCtx(ctx context.Context, key string, members ...string) (
	[]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) GeoRadius(key string, longitude, latitude float64, query *GeoRadiusQuery) (
	[]GeoLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) GeoRadiusCtx(ctx context.Context, key string, longitude, latitude float64,
	query *GeoRadiusQuery) ([]GeoLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) GeoRadiusByMember(key, member string, query *GeoRadiusQuery) ([]GeoLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) GeoRadiusByMemberCtx(ctx context.Context, key, member string,
	query *GeoRadiusQuery) ([]GeoLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) GeoPos(key string, members ...string) ([]*GeoPos, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) GeoPosCtx(ctx context.Context, key string, members ...string) (
	[]*GeoPos, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Get(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *Redis) GetCtx(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) GetBit(key string, offset int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) GetBitCtx(ctx context.Context, key string, offset int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) GetDel(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *Redis) GetDelCtx(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) GetEx(key string, seconds int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) GetExCtx(ctx context.Context, key string, seconds int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) GetSet(key, value string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) GetSetCtx(ctx context.Context, key, value string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) Hdel(key string, fields ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) HdelCtx(ctx context.Context, key string, fields ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) Hexists(key, field string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) HexistsCtx(ctx context.Context, key, field string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) Hget(key, field string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *Redis) HgetCtx(ctx context.Context, key, field string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) Hgetall(key string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) HgetallCtx(ctx context.Context, key string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Hincrby(key, field string, increment int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) HincrbyCtx(ctx context.Context, key, field string, increment int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) HincrbyFloat(key, field string, increment float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) HincrbyFloatCtx(ctx context.Context, key, field string, increment float64) (
	float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Hkeys(key string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Redis) HkeysCtx(ctx context.Context, key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Hlen(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) HlenCtx(ctx context.Context, key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Hmget(key string, fields ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) HmgetCtx(ctx context.Context, key string, fields ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Hset(key, field, value string) error { _ = "STUB: not implemented"; return nil }

func (s *Redis) HsetCtx(ctx context.Context, key, field, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Redis) Hsetnx(key, field, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) HsetnxCtx(ctx context.Context, key, field, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) Hmset(key string, fieldsAndValues map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Redis) HmsetCtx(ctx context.Context, key string, fieldsAndValues map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Redis) Hscan(key string, cursor uint64, match string, count int64) (
	[]string, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *Redis) HscanCtx(ctx context.Context, key string, cursor uint64, match string, count int64) (
	[]string, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *Redis) Hvals(key string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Redis) HvalsCtx(ctx context.Context, key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Incr(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) IncrCtx(ctx context.Context, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Incrby(key string, increment int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) IncrbyCtx(ctx context.Context, key string, increment int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) IncrbyFloat(key string, increment float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) IncrbyFloatCtx(ctx context.Context, key string, increment float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Keys(pattern string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Redis) KeysCtx(ctx context.Context, pattern string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Llen(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) LlenCtx(ctx context.Context, key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Lindex(key string, index int64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) LindexCtx(ctx context.Context, key string, index int64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) Lpop(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *Redis) LpopCtx(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) LpopCount(key string, count int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) LpopCountCtx(ctx context.Context, key string, count int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Lpush(key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) LpushCtx(ctx context.Context, key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Lrange(key string, start, stop int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) LrangeCtx(ctx context.Context, key string, start, stop int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Lrem(key string, count int, value string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) LremCtx(ctx context.Context, key string, count int, value string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Ltrim(key string, start, stop int64) error { _ = "STUB: not implemented"; return nil }

func (s *Redis) LtrimCtx(ctx context.Context, key string, start, stop int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Redis) Mget(keys ...string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Redis) MgetCtx(ctx context.Context, keys ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Mset(fieldsAndValues ...any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) MsetCtx(ctx context.Context, fieldsAndValues ...any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) Persist(key string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s *Redis) PersistCtx(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) Pfadd(key string, values ...any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) PfaddCtx(ctx context.Context, key string, values ...any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) Pfcount(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) PfcountCtx(ctx context.Context, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Pfmerge(dest string, keys ...string) error { _ = "STUB: not implemented"; return nil }

func (s *Redis) PfmergeCtx(ctx context.Context, dest string, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Redis) Ping() bool { _ = "STUB: not implemented"; return false }

func (s *Redis) PingCtx(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (s *Redis) Pipelined(fn func(Pipeliner) error) error { _ = "STUB: not implemented"; return nil }

func (s *Redis) PipelinedCtx(ctx context.Context, fn func(Pipeliner) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Redis) Publish(channel string, message interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) PublishCtx(ctx context.Context, channel string, message interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Rpop(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *Redis) RpopCtx(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) RpopCount(key string, count int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) RpopCountCtx(ctx context.Context, key string, count int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Rpush(key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) RpushCtx(ctx context.Context, key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) RPopLPush(source string, destination string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) RPopLPushCtx(ctx context.Context, source string, destination string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) Sadd(key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) SaddCtx(ctx context.Context, key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *Redis) ScanCtx(ctx context.Context, cursor uint64, match string, count int64) (
	[]string, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *Redis) SetBit(key string, offset int64, value int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) SetBitCtx(ctx context.Context, key string, offset int64, value int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Sscan(key string, cursor uint64, match string, count int64) (
	[]string, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *Redis) SscanCtx(ctx context.Context, key string, cursor uint64, match string, count int64) (
	[]string, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *Redis) Scard(key string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) ScardCtx(ctx context.Context, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) ScriptLoad(script string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) ScriptLoadCtx(ctx context.Context, script string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) ScriptRun(script *Script, keys []string, args ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *Redis) ScriptRunCtx(ctx context.Context, script *Script, keys []string,
	args ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (s *Redis) Set(key, value string) error { _ = "STUB: not implemented"; return nil }

func (s *Redis) SetCtx(ctx context.Context, key, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Redis) Setex(key, value string, seconds int) error { _ = "STUB: not implemented"; return nil }

func (s *Redis) SetexCtx(ctx context.Context, key, value string, seconds int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Redis) Setnx(key, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) SetnxCtx(ctx context.Context, key, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) SetnxEx(key, value string, seconds int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) SetnxExCtx(ctx context.Context, key, value string, seconds int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) Sismember(key string, value any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) SismemberCtx(ctx context.Context, key string, value any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) Smembers(key string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Redis) SmembersCtx(ctx context.Context, key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Spop(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *Redis) SpopCtx(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) Srandmember(key string, count int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) SrandmemberCtx(ctx context.Context, key string, count int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Srem(key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) SremCtx(ctx context.Context, key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) String() string { _ = "STUB: not implemented"; return "" }

func (s *Redis) Sunion(keys ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) SunionCtx(ctx context.Context, keys ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Sunionstore(destination string, keys ...string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) SunionstoreCtx(ctx context.Context, destination string, keys ...string) (
	int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Sdiff(keys ...string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Redis) SdiffCtx(ctx context.Context, keys ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Sdiffstore(destination string, keys ...string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) SdiffstoreCtx(ctx context.Context, destination string, keys ...string) (
	int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Sinter(keys ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) SinterCtx(ctx context.Context, keys ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Sinterstore(destination string, keys ...string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) SinterstoreCtx(ctx context.Context, destination string, keys ...string) (
	int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Ttl(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) TtlCtx(ctx context.Context, key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) TxPipeline() (pipe Pipeliner, err error) {
	_ = "STUB: not implemented"
	return *new(Pipeliner), nil
}

func (s *Redis) Unlink(keys ...string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) UnlinkCtx(ctx context.Context, keys ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) XAck(stream string, group string, ids ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) XAckCtx(ctx context.Context, stream string, group string, ids ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) XAdd(stream string, noMkStream bool, id string, values any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) XAddCtx(ctx context.Context, stream string, noMkStream bool, id string, values any) (
	string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) XGroupCreateMkStream(stream string, group string, start string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) XGroupCreateMkStreamCtx(ctx context.Context, stream string, group string,
	start string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) XGroupCreate(stream string, group string, start string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) XGroupCreateCtx(ctx context.Context, stream string, group string, start string) (
	string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) XGroupSetID(stream, group, start string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) XGroupSetIDCtx(ctx context.Context, stream, group, start string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Redis) XInfoConsumers(stream string, group string) ([]red.XInfoConsumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) XInfoConsumersCtx(ctx context.Context, stream string, group string) (
	[]red.XInfoConsumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) XInfoGroups(stream string) ([]red.XInfoGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) XInfoGroupsCtx(ctx context.Context, stream string) ([]red.XInfoGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) XInfoStream(stream string) (*red.XInfoStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) XInfoStreamCtx(ctx context.Context, stream string) (*red.XInfoStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) XReadGroup(node RedisNode, group string, consumerId string, count int64,
	block time.Duration, noAck bool, streams ...string) ([]red.XStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) XReadGroupCtx(ctx context.Context, node RedisNode, group string, consumerId string,
	count int64, block time.Duration, noAck bool, streams ...string) ([]red.XStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Zadd(key string, score int64, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) ZaddCtx(ctx context.Context, key string, score int64, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) ZaddFloat(key string, score float64, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) ZaddFloatCtx(ctx context.Context, key string, score float64, value string) (
	bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) Zaddnx(key string, score int64, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) ZaddnxCtx(ctx context.Context, key string, score int64, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) ZaddnxFloat(key string, score float64, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) ZaddnxFloatCtx(ctx context.Context, key string, score float64, value string) (
	bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Redis) Zadds(key string, ps ...Pair) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) ZaddsCtx(ctx context.Context, key string, ps ...Pair) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Zcard(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) ZcardCtx(ctx context.Context, key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Zcount(key string, start, stop int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) ZcountCtx(ctx context.Context, key string, start, stop int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Zincrby(key string, increment int64, field string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) ZincrbyCtx(ctx context.Context, key string, increment int64, field string) (
	int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Zscore(key, value string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) ZscoreCtx(ctx context.Context, key, value string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) ZscoreByFloat(key, value string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) ZscoreByFloatCtx(ctx context.Context, key, value string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Zscan(key string, cursor uint64, match string, count int64) (
	[]string, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *Redis) ZscanCtx(ctx context.Context, key string, cursor uint64, match string, count int64) (
	[]string, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *Redis) Zrank(key, field string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Redis) ZrankCtx(ctx context.Context, key, field string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Zrem(key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) ZremCtx(ctx context.Context, key string, values ...any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Zremrangebyscore(key string, start, stop int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) ZremrangebyscoreCtx(ctx context.Context, key string, start, stop int64) (
	int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Zremrangebyrank(key string, start, stop int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) ZremrangebyrankCtx(ctx context.Context, key string, start, stop int64) (
	int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Zrange(key string, start, stop int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangeCtx(ctx context.Context, key string, start, stop int64) (
	[]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangeWithScores(key string, start, stop int64) ([]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangeWithScoresCtx(ctx context.Context, key string, start, stop int64) (
	[]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangeWithScoresByFloat(key string, start, stop int64) ([]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangeWithScoresByFloatCtx(ctx context.Context, key string, start, stop int64) (
	[]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZRevRangeWithScores(key string, start, stop int64) ([]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangeWithScores(key string, start, stop int64) ([]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZRevRangeWithScoresCtx(ctx context.Context, key string, start, stop int64) (
	[]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangeWithScoresCtx(ctx context.Context, key string, start, stop int64) (
	[]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZRevRangeWithScoresByFloat(key string, start, stop int64) ([]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangeWithScoresByFloat(key string, start, stop int64) ([]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZRevRangeWithScoresByFloatCtx(ctx context.Context, key string, start, stop int64) (
	[]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangeWithScoresByFloatCtx(ctx context.Context, key string, start, stop int64) (
	[]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangebyscoreWithScores(key string, start, stop int64) ([]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangebyscoreWithScoresCtx(ctx context.Context, key string, start, stop int64) (
	[]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangebyscoreWithScoresByFloat(key string, start, stop float64) (
	[]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangebyscoreWithScoresByFloatCtx(ctx context.Context, key string, start, stop float64) (
	[]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangebyscoreWithScoresAndLimit(key string, start, stop int64,
	page, size int) ([]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangebyscoreWithScoresAndLimitCtx(ctx context.Context, key string, start,
	stop int64, page, size int) ([]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangebyscoreWithScoresByFloatAndLimit(key string, start, stop float64,
	page, size int) ([]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrangebyscoreWithScoresByFloatAndLimitCtx(ctx context.Context, key string, start,
	stop float64, page, size int) ([]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Zrevrange(key string, start, stop int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangeCtx(ctx context.Context, key string, start, stop int64) (
	[]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangebyscoreWithScores(key string, start, stop int64) ([]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangebyscoreWithScoresCtx(ctx context.Context, key string, start, stop int64) (
	[]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangebyscoreWithScoresByFloat(key string, start, stop float64) (
	[]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangebyscoreWithScoresByFloatCtx(ctx context.Context, key string,
	start, stop float64) ([]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangebyscoreWithScoresAndLimit(key string, start, stop int64,
	page, size int) ([]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangebyscoreWithScoresAndLimitCtx(ctx context.Context, key string,
	start, stop int64, page, size int) ([]Pair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangebyscoreWithScoresByFloatAndLimit(key string, start, stop float64,
	page, size int) ([]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) ZrevrangebyscoreWithScoresByFloatAndLimitCtx(ctx context.Context, key string,
	start, stop float64, page, size int) ([]FloatPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Redis) Zrevrank(key, field string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) ZrevrankCtx(ctx context.Context, key, field string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) Zunionstore(dest string, store *ZStore) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) ZunionstoreCtx(ctx context.Context, dest string, store *ZStore) (
	int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Redis) checkConnection(pingTimeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Redis) maintNotificationsConfig() *maintnotifications.Config {
	_ = "STUB: not implemented"
	return nil
}

func Cluster() Option { _ = "STUB: not implemented"; return *new(Option) }

func SetSlowThreshold(threshold time.Duration) { _ = "STUB: not implemented"; return }

func WithHook(hook Hook) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPass(pass string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTLS() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithUser(user string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProtocol(protocol int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithIdentity() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaintNotifications(mode string) Option { _ = "STUB: not implemented"; return *new(Option) }

func acceptable(err error) bool { _ = "STUB: not implemented"; return false }

func getRedis(r *Redis) (RedisNode, error) { _ = "STUB: not implemented"; return *new(RedisNode), nil }

func newRedis(addr string, opts ...Option) *Redis { _ = "STUB: not implemented"; return nil }

func toPairs(vals []red.Z) []Pair { _ = "STUB: not implemented"; return nil }

func toFloatPairs(vals []red.Z) []FloatPair { _ = "STUB: not implemented"; return nil }

func toStrings(vals []any) []string { _ = "STUB: not implemented"; return nil }
