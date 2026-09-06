//go:generate mockgen -package mon -destination collectioninserter_mock.go -source bulkinserter.go collectionInserter
package mon

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/executors"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	flushInterval = time.Second
	maxBulkRows   = 1000
)

type (
	ResultHandler func(*mongo.InsertManyResult, error)

	BulkInserter struct {
		executor *executors.PeriodicalExecutor
		inserter *dbInserter
	}
)

func NewBulkInserter(coll Collection, interval ...time.Duration) (*BulkInserter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bi *BulkInserter) Flush() { _ = "STUB: not implemented"; return }

func (bi *BulkInserter) Insert(doc any) { _ = "STUB: not implemented"; return }

func (bi *BulkInserter) SetResultHandler(handler ResultHandler) { _ = "STUB: not implemented"; return }

type collectionInserter interface {
	InsertMany(
		ctx context.Context,
		documents interface{},
		opts ...options.Lister[options.InsertManyOptions],
	) (*mongo.InsertManyResult, error)
}

type dbInserter struct {
	collection    collectionInserter
	documents     []any
	resultHandler ResultHandler
}

func (in *dbInserter) AddTask(doc any) bool { _ = "STUB: not implemented"; return false }

func (in *dbInserter) Execute(objs any) { _ = "STUB: not implemented"; return }

func (in *dbInserter) RemoveAll() any { _ = "STUB: not implemented"; return *new(any) }
