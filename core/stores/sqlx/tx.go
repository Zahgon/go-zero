package sqlx

import (
	"context"
	"database/sql"
)

type (
	beginnable func(*sql.DB) (trans, error)

	trans interface {
		Session
		Commit() error
		Rollback() error
	}

	txConn struct {
		Session
	}

	txSession struct {
		*sql.Tx
	}
)

func (s txConn) RawDB() (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func (s txConn) Transact(_ func(Session) error) error { _ = "STUB: not implemented"; return nil }

func (s txConn) TransactCtx(_ context.Context, _ func(context.Context, Session) error) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSessionFromTx(tx *sql.Tx) Session { _ = "STUB: not implemented"; return *new(Session) }

func (t txSession) Exec(q string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (t txSession) ExecCtx(ctx context.Context, q string, args ...any) (result sql.Result, err error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (t txSession) Prepare(q string) (StmtSession, error) {
	_ = "STUB: not implemented"
	return *new(StmtSession), nil
}

func (t txSession) PrepareCtx(ctx context.Context, q string) (stmtSession StmtSession, err error) {
	_ = "STUB: not implemented"
	return *new(StmtSession), nil
}

func (t txSession) QueryRow(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (t txSession) QueryRowCtx(ctx context.Context, v any, q string, args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (t txSession) QueryRowPartial(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (t txSession) QueryRowPartialCtx(ctx context.Context, v any, q string,
	args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (t txSession) QueryRows(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (t txSession) QueryRowsCtx(ctx context.Context, v any, q string, args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (t txSession) QueryRowsPartial(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (t txSession) QueryRowsPartialCtx(ctx context.Context, v any, q string,
	args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func begin(db *sql.DB) (trans, error) { _ = "STUB: not implemented"; return *new(trans), nil }

func transact(ctx context.Context, db *commonSqlConn, b beginnable,
	fn func(context.Context, Session) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func transactOnConn(ctx context.Context, conn *sql.DB, b beginnable,
	fn func(context.Context, Session) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}
