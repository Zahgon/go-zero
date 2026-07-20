package mon

import (
	context "context"

	options "go.mongodb.org/mongo-driver/v2/mongo/options"
	gomock "go.uber.org/mock/gomock"
)

type MockmonClient struct {
	ctrl     *gomock.Controller
	recorder *MockmonClientMockRecorder
	isgomock struct{}
}

type MockmonClientMockRecorder struct {
	mock *MockmonClient
}

func NewMockmonClient(ctrl *gomock.Controller) *MockmonClient {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonClient) EXPECT() *MockmonClientMockRecorder { _ = "STUB: not implemented"; return nil }

func (m *MockmonClient) StartSession(opts ...options.Lister[options.SessionOptions]) (monSession, error) {
	_ = "STUB: not implemented"
	return *new(monSession), nil
}

func (mr *MockmonClientMockRecorder) StartSession(opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

type MockmonSession struct {
	ctrl     *gomock.Controller
	recorder *MockmonSessionMockRecorder
	isgomock struct{}
}

type MockmonSessionMockRecorder struct {
	mock *MockmonSession
}

func NewMockmonSession(ctrl *gomock.Controller) *MockmonSession {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonSession) EXPECT() *MockmonSessionMockRecorder {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonSession) AbortTransaction(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockmonSessionMockRecorder) AbortTransaction(ctx any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonSession) CommitTransaction(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockmonSessionMockRecorder) CommitTransaction(ctx any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonSession) EndSession(ctx context.Context) { _ = "STUB: not implemented"; return }

func (mr *MockmonSessionMockRecorder) EndSession(ctx any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockmonSession) WithTransaction(ctx context.Context, fn func(context.Context) (any, error), opts ...options.Lister[options.TransactionOptions]) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (mr *MockmonSessionMockRecorder) WithTransaction(ctx, fn any, opts ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}
