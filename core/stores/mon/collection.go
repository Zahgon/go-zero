//go:generate mockgen -package mon -destination collection_mock.go -source collection.go Collection,monCollection
package mon

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/breaker"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	defaultSlowThreshold = time.Millisecond * 500

	spanName         = "mongo"
	duplicateKeyCode = 11000

	aggregate              = "Aggregate"
	bulkWrite              = "BulkWrite"
	countDocuments         = "CountDocuments"
	deleteMany             = "DeleteMany"
	deleteOne              = "DeleteOne"
	distinct               = "Distinct"
	estimatedDocumentCount = "EstimatedDocumentCount"
	find                   = "Find"
	findOne                = "FindOne"
	findOneAndDelete       = "FindOneAndDelete"
	findOneAndReplace      = "FindOneAndReplace"
	findOneAndUpdate       = "FindOneAndUpdate"
	insertMany             = "InsertMany"
	insertOne              = "InsertOne"
	replaceOne             = "ReplaceOne"
	updateByID             = "UpdateByID"
	updateMany             = "UpdateMany"
	updateOne              = "UpdateOne"
)

var ErrNotFound = mongo.ErrNoDocuments

type (
	Collection interface {
		Aggregate(ctx context.Context, pipeline any, opts ...options.Lister[options.AggregateOptions]) (
			*mongo.Cursor, error)

		BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...options.Lister[options.BulkWriteOptions]) (
			*mongo.BulkWriteResult, error)

		Clone(opts ...options.Lister[options.CollectionOptions]) *mongo.Collection

		CountDocuments(ctx context.Context, filter any, opts ...options.Lister[options.CountOptions]) (int64, error)

		Database() *mongo.Database

		DeleteMany(ctx context.Context, filter any, opts ...options.Lister[options.DeleteManyOptions]) (
			*mongo.DeleteResult, error)

		DeleteOne(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (
			*mongo.DeleteResult, error)

		Distinct(ctx context.Context, fieldName string, filter any,
			opts ...options.Lister[options.DistinctOptions]) (*mongo.DistinctResult, error)

		Drop(ctx context.Context, opts ...options.Lister[options.DropCollectionOptions]) error

		EstimatedDocumentCount(ctx context.Context, opts ...options.Lister[options.EstimatedDocumentCountOptions]) (int64, error)

		Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error)

		FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) (
			*mongo.SingleResult, error)

		FindOneAndDelete(ctx context.Context, filter any, opts ...options.Lister[options.FindOneAndDeleteOptions]) (
			*mongo.SingleResult, error)

		FindOneAndReplace(ctx context.Context, filter, replacement any,
			opts ...options.Lister[options.FindOneAndReplaceOptions]) (*mongo.SingleResult, error)

		FindOneAndUpdate(ctx context.Context, filter, update any,
			opts ...options.Lister[options.FindOneAndUpdateOptions]) (*mongo.SingleResult, error)

		Indexes() mongo.IndexView

		InsertMany(ctx context.Context, documents []any, opts ...options.Lister[options.InsertManyOptions]) (
			*mongo.InsertManyResult, error)

		InsertOne(ctx context.Context, document any, opts ...options.Lister[options.InsertOneOptions]) (
			*mongo.InsertOneResult, error)

		ReplaceOne(ctx context.Context, filter, replacement any,
			opts ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error)

		UpdateByID(ctx context.Context, id, update any,
			opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error)

		UpdateMany(ctx context.Context, filter, update any,
			opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error)

		UpdateOne(ctx context.Context, filter, update any,
			opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error)

		Watch(ctx context.Context, pipeline any, opts ...options.Lister[options.ChangeStreamOptions]) (
			*mongo.ChangeStream, error)
	}

	decoratedCollection struct {
		Collection monCollection
		name       string
		brk        breaker.Breaker
	}

	keepablePromise struct {
		promise breaker.Promise
		log     func(error)
	}
)

func newCollection(collection *mongo.Collection, brk breaker.Breaker) Collection {
	_ = "STUB: not implemented"
	return *new(Collection)
}

func (c *decoratedCollection) Aggregate(ctx context.Context, pipeline any,
	opts ...options.Lister[options.AggregateOptions]) (cur *mongo.Cursor, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) BulkWrite(ctx context.Context, models []mongo.WriteModel,
	opts ...options.Lister[options.BulkWriteOptions]) (res *mongo.BulkWriteResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) Clone(opts ...options.Lister[options.CollectionOptions]) *mongo.Collection {
	_ = "STUB: not implemented"
	return nil
}

func (c *decoratedCollection) CountDocuments(ctx context.Context, filter any,
	opts ...options.Lister[options.CountOptions]) (count int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *decoratedCollection) Database() *mongo.Database { _ = "STUB: not implemented"; return nil }

func (c *decoratedCollection) DeleteMany(ctx context.Context, filter any,
	opts ...options.Lister[options.DeleteManyOptions]) (res *mongo.DeleteResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) DeleteOne(ctx context.Context, filter any,
	opts ...options.Lister[options.DeleteOneOptions]) (res *mongo.DeleteResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) Distinct(ctx context.Context, fieldName string, filter any,
	opts ...options.Lister[options.DistinctOptions]) (res *mongo.DistinctResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) Drop(ctx context.Context, opts ...options.Lister[options.DropCollectionOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *decoratedCollection) EstimatedDocumentCount(ctx context.Context,
	opts ...options.Lister[options.EstimatedDocumentCountOptions]) (val int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *decoratedCollection) Find(ctx context.Context, filter any,
	opts ...options.Lister[options.FindOptions]) (cur *mongo.Cursor, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) FindOne(ctx context.Context, filter any,
	opts ...options.Lister[options.FindOneOptions]) (res *mongo.SingleResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) FindOneAndDelete(ctx context.Context, filter any,
	opts ...options.Lister[options.FindOneAndDeleteOptions]) (res *mongo.SingleResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) FindOneAndReplace(ctx context.Context, filter any,
	replacement any, opts ...options.Lister[options.FindOneAndReplaceOptions]) (
	res *mongo.SingleResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) FindOneAndUpdate(ctx context.Context, filter, update any,
	opts ...options.Lister[options.FindOneAndUpdateOptions]) (res *mongo.SingleResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) Indexes() mongo.IndexView {
	_ = "STUB: not implemented"
	return *new(mongo.IndexView)
}

func (c *decoratedCollection) InsertMany(ctx context.Context, documents []any,
	opts ...options.Lister[options.InsertManyOptions]) (res *mongo.InsertManyResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) InsertOne(ctx context.Context, document any,
	opts ...options.Lister[options.InsertOneOptions]) (res *mongo.InsertOneResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) ReplaceOne(ctx context.Context, filter, replacement any,
	opts ...options.Lister[options.ReplaceOptions]) (res *mongo.UpdateResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) UpdateByID(ctx context.Context, id, update any,
	opts ...options.Lister[options.UpdateOneOptions]) (res *mongo.UpdateResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) UpdateMany(ctx context.Context, filter, update any,
	opts ...options.Lister[options.UpdateManyOptions]) (res *mongo.UpdateResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) UpdateOne(ctx context.Context, filter, update any,
	opts ...options.Lister[options.UpdateOneOptions]) (res *mongo.UpdateResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) Watch(ctx context.Context, pipeline any, opts ...options.Lister[options.ChangeStreamOptions]) (
	*mongo.ChangeStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *decoratedCollection) logDuration(ctx context.Context, method string,
	startTime time.Duration, err error, docs ...any) {
	_ = "STUB: not implemented"
	return
}

func (c *decoratedCollection) logDurationSimple(ctx context.Context, method string,
	startTime time.Duration, err error) {
	_ = "STUB: not implemented"
	return
}

func (p keepablePromise) accept(err error) error { _ = "STUB: not implemented"; return nil }

func (p keepablePromise) keep(err error) error { _ = "STUB: not implemented"; return nil }

func acceptable(err error) bool { _ = "STUB: not implemented"; return false }

func isDupKeyError(err error) bool { _ = "STUB: not implemented"; return false }

type monCollection interface {
	Aggregate(ctx context.Context, pipeline any, opts ...options.Lister[options.AggregateOptions]) (
		*mongo.Cursor, error)

	BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...options.Lister[options.BulkWriteOptions]) (
		*mongo.BulkWriteResult, error)

	Clone(opts ...options.Lister[options.CollectionOptions]) *mongo.Collection

	CountDocuments(ctx context.Context, filter any, opts ...options.Lister[options.CountOptions]) (int64, error)

	Database() *mongo.Database

	DeleteMany(ctx context.Context, filter any, opts ...options.Lister[options.DeleteManyOptions]) (
		*mongo.DeleteResult, error)

	DeleteOne(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (
		*mongo.DeleteResult, error)

	Distinct(ctx context.Context, fieldName string, filter any,
		opts ...options.Lister[options.DistinctOptions]) *mongo.DistinctResult

	Drop(ctx context.Context, opts ...options.Lister[options.DropCollectionOptions]) error

	EstimatedDocumentCount(ctx context.Context, opts ...options.Lister[options.EstimatedDocumentCountOptions]) (int64, error)

	Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error)

	FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult

	FindOneAndDelete(ctx context.Context, filter any, opts ...options.Lister[options.FindOneAndDeleteOptions]) *mongo.SingleResult

	FindOneAndReplace(ctx context.Context, filter, replacement any,
		opts ...options.Lister[options.FindOneAndReplaceOptions]) *mongo.SingleResult

	FindOneAndUpdate(ctx context.Context, filter, update any,
		opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult

	Indexes() mongo.IndexView

	InsertMany(ctx context.Context, documents interface{}, opts ...options.Lister[options.InsertManyOptions]) (*mongo.InsertManyResult, error)

	InsertOne(ctx context.Context, document any, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error)

	ReplaceOne(ctx context.Context, filter, replacement any,
		opts ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error)

	UpdateByID(ctx context.Context, id, update any,
		opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error)

	UpdateMany(ctx context.Context, filter, update any,
		opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error)

	UpdateOne(ctx context.Context, filter, update any,
		opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error)

	Watch(ctx context.Context, pipeline any, opts ...options.Lister[options.ChangeStreamOptions]) (
		*mongo.ChangeStream, error)
}
