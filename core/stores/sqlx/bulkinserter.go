package sqlx

import (
	"database/sql"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/executors"
)

const (
	flushInterval = time.Second
	maxBulkRows   = 1000
	valuesKeyword = "values"
)

var emptyBulkStmt bulkStmt

type (
	ResultHandler func(sql.Result, error)

	BulkInserter struct {
		executor *executors.PeriodicalExecutor
		inserter *dbInserter
		stmt     bulkStmt
		lock     sync.RWMutex
	}

	bulkStmt struct {
		prefix      string
		valueFormat string
		suffix      string
	}
)

func NewBulkInserter(sqlConn SqlConn, stmt string) (*BulkInserter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bi *BulkInserter) Flush() { _ = "STUB: not implemented"; return }

func (bi *BulkInserter) Insert(args ...any) error { _ = "STUB: not implemented"; return nil }

func (bi *BulkInserter) SetResultHandler(handler ResultHandler) { _ = "STUB: not implemented"; return }

func (bi *BulkInserter) UpdateOrDelete(fn func()) { _ = "STUB: not implemented"; return }

func (bi *BulkInserter) UpdateStmt(stmt string) error { _ = "STUB: not implemented"; return nil }

type dbInserter struct {
	sqlConn       SqlConn
	stmt          bulkStmt
	values        []string
	resultHandler ResultHandler
}

func (in *dbInserter) AddTask(task any) bool { _ = "STUB: not implemented"; return false }

func (in *dbInserter) Execute(bulk any) { _ = "STUB: not implemented"; return }

func (in *dbInserter) RemoveAll() any { _ = "STUB: not implemented"; return *new(any) }

func parseInsertStmt(stmt string) (bulkStmt, error) {
	_ = "STUB: not implemented"
	return *new(bulkStmt), nil
}
