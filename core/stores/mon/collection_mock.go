package mon

import (
	context "context"

	mongo "go.mongodb.org/mongo-driver/v2/mongo"
	options "go.mongodb.org/mongo-driver/v2/mongo/options"
	gomock "go.uber.org/mock/gomock"
)

type MockCollection struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionMockRecorder
	isgomock struct{}
}

type MockCollectionMockRecorder struct {
	mock *MockCollection
}

func NewMockCollection(ctrl *gomock.Controller) *MockCollection {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) EXPECT() *MockCollectionMockRecorder {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) Aggregate(ctx context.Context, pipeline any, opts ...options.Lister[options.AggregateOptions]) (*mongo.Cursor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) Aggregate(ctx, pipeline any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...options.Lister[options.BulkWriteOptions]) (*mongo.BulkWriteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) BulkWrite(ctx, models any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) Clone(opts ...options.Lister[options.CollectionOptions]) *mongo.Collection {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockCollectionMockRecorder) Clone(opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) CountDocuments(ctx context.Context, filter any, opts ...options.Lister[options.CountOptions]) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mr *MockCollectionMockRecorder) CountDocuments(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) Database() *mongo.Database { _ = "STUB: not implemented"; return nil }

func (mr *MockCollectionMockRecorder) Database() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) DeleteMany(ctx context.Context, filter any, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) DeleteMany(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) DeleteOne(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) Distinct(ctx context.Context, fieldName string, filter any, opts ...options.Lister[options.DistinctOptions]) (*mongo.DistinctResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) Distinct(ctx, fieldName, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) Drop(ctx context.Context, opts ...options.Lister[options.DropCollectionOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockCollectionMockRecorder) Drop(ctx any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) EstimatedDocumentCount(ctx context.Context, opts ...options.Lister[options.EstimatedDocumentCountOptions]) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mr *MockCollectionMockRecorder) EstimatedDocumentCount(ctx any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) Find(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) (*mongo.SingleResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) FindOne(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) FindOneAndDelete(ctx context.Context, filter any, opts ...options.Lister[options.FindOneAndDeleteOptions]) (*mongo.SingleResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) FindOneAndDelete(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) FindOneAndReplace(ctx context.Context, filter, replacement any, opts ...options.Lister[options.FindOneAndReplaceOptions]) (*mongo.SingleResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) FindOneAndReplace(ctx, filter, replacement any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) FindOneAndUpdate(ctx context.Context, filter, update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) (*mongo.SingleResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) FindOneAndUpdate(ctx, filter, update any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) Indexes() mongo.IndexView {
	_ = "STUB: not implemented"
	return *new(mongo.IndexView)
}

func (mr *MockCollectionMockRecorder) Indexes() *gomock.Call { _ = "STUB: not implemented"; return nil }

func (m *MockCollection) InsertMany(ctx context.Context, documents []any, opts ...options.Lister[options.InsertManyOptions]) (*mongo.InsertManyResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) InsertMany(ctx, documents any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) InsertOne(ctx context.Context, document any, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) InsertOne(ctx, document any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) ReplaceOne(ctx context.Context, filter, replacement any, opts ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) ReplaceOne(ctx, filter, replacement any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) UpdateByID(ctx context.Context, id, update any, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) UpdateByID(ctx, id, update any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) UpdateMany(ctx context.Context, filter, update any, opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) UpdateMany(ctx, filter, update any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter, update any, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) UpdateOne(ctx, filter, update any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockCollection) Watch(ctx context.Context, pipeline any, opts ...options.Lister[options.ChangeStreamOptions]) (*mongo.ChangeStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockCollectionMockRecorder) Watch(ctx, pipeline any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

type MockmonCollection struct {
	ctrl     *gomock.Controller
	recorder *MockmonCollectionMockRecorder
	isgomock struct{}
}

type MockmonCollectionMockRecorder struct {
	mock *MockmonCollection
}

func NewMockmonCollection(ctrl *gomock.Controller) *MockmonCollection {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) EXPECT() *MockmonCollectionMockRecorder {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) Aggregate(ctx context.Context, pipeline any, opts ...options.Lister[options.AggregateOptions]) (*mongo.Cursor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) Aggregate(ctx, pipeline any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...options.Lister[options.BulkWriteOptions]) (*mongo.BulkWriteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) BulkWrite(ctx, models any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) Clone(opts ...options.Lister[options.CollectionOptions]) *mongo.Collection {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockmonCollectionMockRecorder) Clone(opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) CountDocuments(ctx context.Context, filter any, opts ...options.Lister[options.CountOptions]) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mr *MockmonCollectionMockRecorder) CountDocuments(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) Database() *mongo.Database { _ = "STUB: not implemented"; return nil }

func (mr *MockmonCollectionMockRecorder) Database() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) DeleteMany(ctx context.Context, filter any, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) DeleteMany(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) DeleteOne(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) DeleteOne(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) Distinct(ctx context.Context, fieldName string, filter any, opts ...options.Lister[options.DistinctOptions]) *mongo.DistinctResult {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockmonCollectionMockRecorder) Distinct(ctx, fieldName, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) Drop(ctx context.Context, opts ...options.Lister[options.DropCollectionOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockmonCollectionMockRecorder) Drop(ctx any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) EstimatedDocumentCount(ctx context.Context, opts ...options.Lister[options.EstimatedDocumentCountOptions]) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mr *MockmonCollectionMockRecorder) EstimatedDocumentCount(ctx any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) Find(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockmonCollectionMockRecorder) FindOne(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) FindOneAndDelete(ctx context.Context, filter any, opts ...options.Lister[options.FindOneAndDeleteOptions]) *mongo.SingleResult {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockmonCollectionMockRecorder) FindOneAndDelete(ctx, filter any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) FindOneAndReplace(ctx context.Context, filter, replacement any, opts ...options.Lister[options.FindOneAndReplaceOptions]) *mongo.SingleResult {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockmonCollectionMockRecorder) FindOneAndReplace(ctx, filter, replacement any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) FindOneAndUpdate(ctx context.Context, filter, update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockmonCollectionMockRecorder) FindOneAndUpdate(ctx, filter, update any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) Indexes() mongo.IndexView {
	_ = "STUB: not implemented"
	return *new(mongo.IndexView)
}

func (mr *MockmonCollectionMockRecorder) Indexes() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) InsertMany(ctx context.Context, documents any, opts ...options.Lister[options.InsertManyOptions]) (*mongo.InsertManyResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) InsertMany(ctx, documents any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) InsertOne(ctx context.Context, document any, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) InsertOne(ctx, document any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) ReplaceOne(ctx context.Context, filter, replacement any, opts ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) ReplaceOne(ctx, filter, replacement any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) UpdateByID(ctx context.Context, id, update any, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) UpdateByID(ctx, id, update any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) UpdateMany(ctx context.Context, filter, update any, opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) UpdateMany(ctx, filter, update any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) UpdateOne(ctx context.Context, filter, update any, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) UpdateOne(ctx, filter, update any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonCollection) Watch(ctx context.Context, pipeline any, opts ...options.Lister[options.ChangeStreamOptions]) (*mongo.ChangeStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockmonCollectionMockRecorder) Watch(ctx, pipeline any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}
