package mon

import (
	context "context"

	mongo "go.mongodb.org/mongo-driver/v2/mongo"
	options "go.mongodb.org/mongo-driver/v2/mongo/options"
	gomock "go.uber.org/mock/gomock"
)

type MockcollectionInserter struct {
	ctrl     *gomock.Controller
	recorder *MockcollectionInserterMockRecorder
	isgomock struct{}
}

type MockcollectionInserterMockRecorder struct {
	mock *MockcollectionInserter
}

func NewMockcollectionInserter(ctrl *gomock.Controller) *MockcollectionInserter {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockcollectionInserter) EXPECT() *MockcollectionInserterMockRecorder {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockcollectionInserter) InsertMany(ctx context.Context, documents any, opts ...options.Lister[options.InsertManyOptions]) (*mongo.InsertManyResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockcollectionInserterMockRecorder) InsertMany(ctx, documents any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}
