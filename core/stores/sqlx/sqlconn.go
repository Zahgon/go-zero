package sqlx

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/breaker"
)

const spanName = "sql"

type (
	Session interface {
		Exec(query string, args ...any) (sql.Result, error)
		ExecCtx(ctx context.Context, query string, args ...any) (sql.Result, error)
		Prepare(query string) (StmtSession, error)
		PrepareCtx(ctx context.Context, query string) (StmtSession, error)
		QueryRow(v any, query string, args ...any) error
		QueryRowCtx(ctx context.Context, v any, query string, args ...any) error
		QueryRowPartial(v any, query string, args ...any) error
		QueryRowPartialCtx(ctx context.Context, v any, query string, args ...any) error
		QueryRows(v any, query string, args ...any) error
		QueryRowsCtx(ctx context.Context, v any, query string, args ...any) error
		QueryRowsPartial(v any, query string, args ...any) error
		QueryRowsPartialCtx(ctx context.Context, v any, query string, args ...any) error
	}

	SqlConn interface {
		Session

		RawDB() (*sql.DB, error)
		Transact(fn func(Session) error) error
		TransactCtx(ctx context.Context, fn func(context.Context, Session) error) error
	}

	SqlOption func(*commonSqlConn)

	commonSqlConn struct {
		connProv connProvider
		onError  func(context.Context, error)
		beginTx  beginnable
		brk      breaker.Breaker
		accept   breaker.Acceptable
		index    uint32
	}

	connProvider func(ctx context.Context) (*sql.DB, error)

	sessionConn interface {
		Exec(query string, args ...any) (sql.Result, error)
		ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
		Query(query string, args ...any) (*sql.Rows, error)
		QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	}
)

func MustNewConn(c SqlConf, opts ...SqlOption) SqlConn {
	_ = "STUB: not implemented"
	return *new(SqlConn)
}

func NewConn(c SqlConf, opts ...SqlOption) (SqlConn, error) {
	_ = "STUB: not implemented"
	return *new(SqlConn), nil
}

func NewSqlConn(driverName, datasource string, opts ...SqlOption) SqlConn {
	_ = "STUB: not implemented"
	return *new(SqlConn)
}

func NewSqlConnFromDB(db *sql.DB, opts ...SqlOption) SqlConn {
	_ = "STUB: not implemented"
	return *new(SqlConn)
}

func NewSqlConnFromSession(session Session) SqlConn {
	_ = "STUB: not implemented"
	return *new(SqlConn)
}

func (db *commonSqlConn) Exec(q string, args ...any) (result sql.Result, err error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (db *commonSqlConn) ExecCtx(ctx context.Context, q string, args ...any) (
	result sql.Result, err error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (db *commonSqlConn) Prepare(query string) (stmt StmtSession, err error) {
	_ = "STUB: not implemented"
	return *new(StmtSession), nil
}

func (db *commonSqlConn) PrepareCtx(ctx context.Context, query string) (stmt StmtSession, err error) {
	_ = "STUB: not implemented"
	return *new(StmtSession), nil
}

func (db *commonSqlConn) QueryRow(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *commonSqlConn) QueryRowCtx(ctx context.Context, v any, q string,
	args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (db *commonSqlConn) QueryRowPartial(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *commonSqlConn) QueryRowPartialCtx(ctx context.Context, v any,
	q string, args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (db *commonSqlConn) QueryRows(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *commonSqlConn) QueryRowsCtx(ctx context.Context, v any, q string,
	args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (db *commonSqlConn) QueryRowsPartial(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *commonSqlConn) QueryRowsPartialCtx(ctx context.Context, v any,
	q string, args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (db *commonSqlConn) RawDB() (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func (db *commonSqlConn) Transact(fn func(Session) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *commonSqlConn) TransactCtx(ctx context.Context, fn func(context.Context, Session) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (db *commonSqlConn) acceptable(err error) bool { _ = "STUB: not implemented"; return false }

func (db *commonSqlConn) queryRows(ctx context.Context, scanner func(*sql.Rows) error,
	q string, args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func getConnProvider(sc *commonSqlConn, driverName, datasource, policy string, replicas []string) connProvider {
	_ = "STUB: not implemented"
	return *new(connProvider)
}

func WithAcceptable(acceptable func(err error) bool) SqlOption {
	_ = "STUB: not implemented"
	return *new(SqlOption)
}
