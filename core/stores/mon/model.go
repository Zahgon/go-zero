//go:generate mockgen -package mon -destination model_mock.go -source model.go monClient monSession
package mon

import (
	"context"

	"github.com/zeromicro/go-zero/core/breaker"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	startSession      = "StartSession"
	abortTransaction  = "AbortTransaction"
	commitTransaction = "CommitTransaction"
	withTransaction   = "WithTransaction"
	endSession        = "EndSession"
)

type (
	Model struct {
		Collection
		name string
		cli  monClient
		brk  breaker.Breaker
		opts []Option
	}

	Session struct {
		session monSession
		name    string
		brk     breaker.Breaker
	}
)

func MustNewModel(uri, db, collection string, opts ...Option) *Model {
	_ = "STUB: not implemented"
	return nil
}

func NewModel(uri, db, collection string, opts ...Option) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newModel(name string, cli *mongo.Client, coll Collection, brk breaker.Breaker,
	opts ...Option) *Model {
	_ = "STUB: not implemented"
	return nil
}

func (m *Model) StartSession(opts ...options.Lister[options.SessionOptions]) (sess *Session, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Model) Aggregate(ctx context.Context, v, pipeline any,
	opts ...options.Lister[options.AggregateOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Model) DeleteMany(ctx context.Context, filter any,
	opts ...options.Lister[options.DeleteManyOptions]) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Model) DeleteOne(ctx context.Context, filter any,
	opts ...options.Lister[options.DeleteOneOptions]) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Model) Find(ctx context.Context, v, filter any,
	opts ...options.Lister[options.FindOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Model) FindOne(ctx context.Context, v, filter any,
	opts ...options.Lister[options.FindOneOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Model) FindOneAndDelete(ctx context.Context, v, filter any,
	opts ...options.Lister[options.FindOneAndDeleteOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Model) FindOneAndReplace(ctx context.Context, v, filter, replacement any,
	opts ...options.Lister[options.FindOneAndReplaceOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Model) FindOneAndUpdate(ctx context.Context, v, filter, update any,
	opts ...options.Lister[options.FindOneAndUpdateOptions]) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Session) AbortTransaction(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w *Session) CommitTransaction(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w *Session) WithTransaction(
	ctx context.Context,
	fn func(sessCtx context.Context) (any, error),
	opts ...options.Lister[options.TransactionOptions],
) (res any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (w *Session) EndSession(ctx context.Context) { _ = "STUB: not implemented"; return }

type (
	monClient interface {
		StartSession(opts ...options.Lister[options.SessionOptions]) (monSession, error)
	}

	monSession interface {
		AbortTransaction(ctx context.Context) error
		CommitTransaction(ctx context.Context) error
		EndSession(ctx context.Context)
		WithTransaction(ctx context.Context, fn func(sessCtx context.Context) (any, error),
			opts ...options.Lister[options.TransactionOptions]) (any, error)
	}
)

type wrappedMonClient struct {
	c *mongo.Client
}

func (m *wrappedMonClient) StartSession(opts ...options.Lister[options.SessionOptions]) (
	monSession, error) {
	_ = "STUB: not implemented"
	return *new(monSession), nil
}
