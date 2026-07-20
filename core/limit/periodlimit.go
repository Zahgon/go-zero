package limit

import (
	"context"
	_ "embed"
	"errors"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	Unknown = iota

	Allowed

	HitQuota

	OverQuota

	internalOverQuota = 0
	internalAllowed   = 1
	internalHitQuota  = 2
)

var (
	ErrUnknownCode = errors.New("unknown status code")

	//go:embed periodscript.lua
	periodLuaScript string
	periodScript    = redis.NewScript(periodLuaScript)
)

type (
	PeriodOption func(l *PeriodLimit)

	PeriodLimit struct {
		period     int
		quota      int
		limitStore *redis.Redis
		keyPrefix  string
		align      bool
	}
)

func NewPeriodLimit(period, quota int, limitStore *redis.Redis, keyPrefix string,
	opts ...PeriodOption) *PeriodLimit {
	_ = "STUB: not implemented"
	return nil
}

func (h *PeriodLimit) Take(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (h *PeriodLimit) TakeCtx(ctx context.Context, key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *PeriodLimit) calcExpireSeconds() int { _ = "STUB: not implemented"; return 0 }

func Align() PeriodOption { _ = "STUB: not implemented"; return *new(PeriodOption) }
