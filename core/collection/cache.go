package collection

import (
	"container/list"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/mathx"
	"github.com/zeromicro/go-zero/core/syncx"
)

const (
	defaultCacheName = "proc"
	slots            = 300
	statInterval     = time.Minute

	expiryDeviation = 0.05
)

var emptyLruCache = emptyLru{}

type (
	CacheOption func(cache *Cache)

	Cache struct {
		name           string
		lock           sync.Mutex
		data           map[string]any
		expire         time.Duration
		timingWheel    *TimingWheel
		lruCache       lru
		barrier        syncx.SingleFlight
		unstableExpiry mathx.Unstable
		stats          *cacheStat
	}
)

func NewCache(expire time.Duration, opts ...CacheOption) (*Cache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cache) Del(key string) { _ = "STUB: not implemented"; return }

func (c *Cache) Get(key string) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

func (c *Cache) Set(key string, value any) { _ = "STUB: not implemented"; return }

func (c *Cache) SetWithExpire(key string, value any, expire time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (c *Cache) Take(key string, fetch func() (any, error)) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (c *Cache) doGet(key string) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

func (c *Cache) onEvict(key string) { _ = "STUB: not implemented"; return }

func (c *Cache) size() int { _ = "STUB: not implemented"; return 0 }

func WithLimit(limit int) CacheOption { _ = "STUB: not implemented"; return *new(CacheOption) }

func WithName(name string) CacheOption { _ = "STUB: not implemented"; return *new(CacheOption) }

type (
	lru interface {
		add(key string)
		remove(key string)
	}

	emptyLru struct{}

	keyLru struct {
		limit    int
		evicts   *list.List
		elements map[string]*list.Element
		onEvict  func(key string)
	}
)

func (elru emptyLru) add(string) { _ = "STUB: not implemented"; return }

func (elru emptyLru) remove(string) { _ = "STUB: not implemented"; return }

func newKeyLru(limit int, onEvict func(key string)) *keyLru { _ = "STUB: not implemented"; return nil }

func (klru *keyLru) add(key string) { _ = "STUB: not implemented"; return }

func (klru *keyLru) remove(key string) { _ = "STUB: not implemented"; return }

func (klru *keyLru) removeOldest() { _ = "STUB: not implemented"; return }

func (klru *keyLru) removeElement(e *list.Element) { _ = "STUB: not implemented"; return }

type cacheStat struct {
	name         string
	hit          uint64
	miss         uint64
	sizeCallback func() int
}

func newCacheStat(name string, sizeCallback func() int) *cacheStat {
	_ = "STUB: not implemented"
	return nil
}

func (cs *cacheStat) IncrementHit() { _ = "STUB: not implemented"; return }

func (cs *cacheStat) IncrementMiss() { _ = "STUB: not implemented"; return }

func (cs *cacheStat) statLoop() { _ = "STUB: not implemented"; return }
