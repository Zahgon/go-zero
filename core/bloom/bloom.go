package bloom

import (
	"context"
	_ "embed"
	"errors"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

const maps = 14

var (
	ErrTooLargeOffset = errors.New("too large offset")

	//go:embed setscript.lua
	setLuaScript string
	setScript    = redis.NewScript(setLuaScript)

	//go:embed testscript.lua
	testLuaScript string
	testScript    = redis.NewScript(testLuaScript)
)

type (
	Filter struct {
		bits   uint
		bitSet bitSetProvider
	}

	bitSetProvider interface {
		check(ctx context.Context, offsets []uint) (bool, error)
		set(ctx context.Context, offsets []uint) error
	}
)

func New(store *redis.Redis, key string, bits uint) *Filter { _ = "STUB: not implemented"; return nil }

func (f *Filter) Add(data []byte) error { _ = "STUB: not implemented"; return nil }

func (f *Filter) AddCtx(ctx context.Context, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Filter) Exists(data []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (f *Filter) ExistsCtx(ctx context.Context, data []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *Filter) getLocations(data []byte) []uint { _ = "STUB: not implemented"; return nil }

type redisBitSet struct {
	store *redis.Redis
	key   string
	bits  uint
}

func newRedisBitSet(store *redis.Redis, key string, bits uint) *redisBitSet {
	_ = "STUB: not implemented"
	return nil
}

func (r *redisBitSet) buildOffsetArgs(offsets []uint) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *redisBitSet) check(ctx context.Context, offsets []uint) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *redisBitSet) del() error { _ = "STUB: not implemented"; return nil }

func (r *redisBitSet) expire(seconds int) error { _ = "STUB: not implemented"; return nil }

func (r *redisBitSet) set(ctx context.Context, offsets []uint) error {
	_ = "STUB: not implemented"
	return nil
}
