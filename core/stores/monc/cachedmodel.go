package monc

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	ErrNotFound = mongo.ErrNoDocuments

	singleFlight = syncx.NewSingleFlight()
	stats        = cache.NewStat("monc")
)

type Model struct {
	*mon.Model
	cache cache.Cache
}

func MustNewModel(uri, db, collection string, c cache.CacheConf, opts ...cache.Option) *Model {
	_ = "STUB: not implemented"
	return nil
}

func MustNewNodeModel(uri, db, collection string, rds *redis.Redis, opts ...cache.Option) *Model {
	_ = "STUB: not implemented"
	return nil
}

func NewModel(uri, db, collection string, conf cache.CacheConf, opts ...cache.Option) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewModelWithCache(uri, db, collection string, c cache.Cache) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewNodeModel(uri, db, collection string, rds *redis.Redis, opts ...cache.Option) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newModel(uri, db, collection string, c cache.Cache) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mm *Model) DelCache(ctx context.Context, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (mm *Model) DeleteOne(ctx context.Context, key string, filter any,
	opts ...options.Lister[options.DeleteOneOptions]) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mm *Model) DeleteOneNoCache(ctx context.Context, filter any,
	opts ...options.Lister[options.DeleteOneOptions]) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mm *Model) FindOne(ctx context.Context, key string, v, filter any,
	opts ...options.Lister[options.FindOneOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (mm *Model) FindOneNoCache(ctx context.Context, v, filter any,
	opts ...options.Lister[options.FindOneOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (mm *Model) FindOneAndDelete(ctx context.Context, key string, v, filter any,
	opts ...options.Lister[options.FindOneAndDeleteOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (mm *Model) FindOneAndDeleteNoCache(ctx context.Context, v, filter any,
	opts ...options.Lister[options.FindOneAndDeleteOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (mm *Model) FindOneAndReplace(ctx context.Context, key string, v, filter any,
	replacement any, opts ...options.Lister[options.FindOneAndReplaceOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (mm *Model) FindOneAndReplaceNoCache(ctx context.Context, v, filter any,
	replacement any, opts ...options.Lister[options.FindOneAndReplaceOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (mm *Model) FindOneAndUpdate(ctx context.Context, key string, v, filter any,
	update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (mm *Model) FindOneAndUpdateNoCache(ctx context.Context, v, filter any,
	update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (mm *Model) GetCache(key string, v any) error { _ = "STUB: not implemented"; return nil }

func (mm *Model) InsertOne(ctx context.Context, key string, document any,
	opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mm *Model) InsertOneNoCache(ctx context.Context, document any,
	opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mm *Model) ReplaceOne(ctx context.Context, key string, filter, replacement any,
	opts ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mm *Model) ReplaceOneNoCache(ctx context.Context, filter, replacement any,
	opts ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mm *Model) SetCache(key string, v any) error { _ = "STUB: not implemented"; return nil }

func (mm *Model) UpdateByID(ctx context.Context, key string, id, update any,
	opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mm *Model) UpdateByIDNoCache(ctx context.Context, id, update any,
	opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mm *Model) UpdateMany(ctx context.Context, keys []string, filter, update any,
	opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mm *Model) UpdateManyNoCache(ctx context.Context, filter, update any,
	opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mm *Model) UpdateOne(ctx context.Context, key string, filter, update any,
	opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mm *Model) UpdateOneNoCache(ctx context.Context, filter, update any,
	opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
