package sqlx

import (
	"context"
	"database/sql"
	"time"

	"github.com/zeromicro/go-zero/core/breaker"
	"github.com/zeromicro/go-zero/core/syncx"
)

const defaultSlowThreshold = time.Millisecond * 500

var (
	slowThreshold = syncx.ForAtomicDuration(defaultSlowThreshold)
	logSql        = syncx.ForAtomicBool(true)
	logSlowSql    = syncx.ForAtomicBool(true)
)

type (
	StmtSession interface {
		Close() error
		Exec(args ...any) (sql.Result, error)
		ExecCtx(ctx context.Context, args ...any) (sql.Result, error)
		QueryRow(v any, args ...any) error
		QueryRowCtx(ctx context.Context, v any, args ...any) error
		QueryRowPartial(v any, args ...any) error
		QueryRowPartialCtx(ctx context.Context, v any, args ...any) error
		QueryRows(v any, args ...any) error
		QueryRowsCtx(ctx context.Context, v any, args ...any) error
		QueryRowsPartial(v any, args ...any) error
		QueryRowsPartialCtx(ctx context.Context, v any, args ...any) error
	}

	statement struct {
		query  string
		stmt   *sql.Stmt
		brk    breaker.Breaker
		accept breaker.Acceptable
	}

	stmtConn interface {
		Exec(args ...any) (sql.Result, error)
		ExecContext(ctx context.Context, args ...any) (sql.Result, error)
		Query(args ...any) (*sql.Rows, error)
		QueryContext(ctx context.Context, args ...any) (*sql.Rows, error)
	}
)

func (s statement) Close() error { _ = "STUB: not implemented"; return nil }

func (s statement) Exec(args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (s statement) ExecCtx(ctx context.Context, args ...any) (result sql.Result, err error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (s statement) QueryRow(v any, args ...any) error { _ = "STUB: not implemented"; return nil }

func (s statement) QueryRowCtx(ctx context.Context, v any, args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s statement) QueryRowPartial(v any, args ...any) error { _ = "STUB: not implemented"; return nil }

func (s statement) QueryRowPartialCtx(ctx context.Context, v any, args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s statement) QueryRows(v any, args ...any) error { _ = "STUB: not implemented"; return nil }

func (s statement) QueryRowsCtx(ctx context.Context, v any, args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s statement) QueryRowsPartial(v any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s statement) QueryRowsPartialCtx(ctx context.Context, v any, args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s statement) queryRows(ctx context.Context, scanFn func(any, rowsScanner) error,
	v any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func DisableLog() { _ = "STUB: not implemented"; return }

func DisableStmtLog() { _ = "STUB: not implemented"; return }

func SetSlowThreshold(threshold time.Duration) { _ = "STUB: not implemented"; return }

func exec(ctx context.Context, conn sessionConn, q string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func execStmt(ctx context.Context, conn stmtConn, q string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func query(ctx context.Context, conn sessionConn, scanner func(*sql.Rows) error,
	q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func queryStmt(ctx context.Context, conn stmtConn, scanner func(*sql.Rows) error,
	q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

type (
	sqlGuard interface {
		start(q string, args ...any) error
		finish(ctx context.Context, err error)
	}

	nilGuard struct{}

	realSqlGuard struct {
		command   string
		stmt      string
		startTime time.Duration
	}
)

func newGuard(command string) sqlGuard { _ = "STUB: not implemented"; return *new(sqlGuard) }

func (n nilGuard) start(_ string, _ ...any) error { _ = "STUB: not implemented"; return nil }

func (n nilGuard) finish(_ context.Context, _ error) { _ = "STUB: not implemented"; return }

func (e *realSqlGuard) finish(ctx context.Context, err error) { _ = "STUB: not implemented"; return }

func (e *realSqlGuard) start(q string, args ...any) error { _ = "STUB: not implemented"; return nil }
