package mocksql

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type (
	MockConn struct {
		db *sql.DB
	}

	statement struct {
		stmt *sql.Stmt
	}
)

func NewMockConn(db *sql.DB) *MockConn { _ = "STUB: not implemented"; return nil }

func (conn *MockConn) Exec(query string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (conn *MockConn) ExecCtx(_ context.Context, query string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (conn *MockConn) Prepare(query string) (sqlx.StmtSession, error) {
	_ = "STUB: not implemented"
	return *new(sqlx.StmtSession), nil
}

func (conn *MockConn) PrepareCtx(_ context.Context, query string) (sqlx.StmtSession, error) {
	_ = "STUB: not implemented"
	return *new(sqlx.StmtSession), nil
}

func (conn *MockConn) QueryRow(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (conn *MockConn) QueryRowCtx(_ context.Context, v any, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (conn *MockConn) QueryRowPartial(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (conn *MockConn) QueryRowPartialCtx(_ context.Context, v any, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (conn *MockConn) QueryRows(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (conn *MockConn) QueryRowsCtx(_ context.Context, v any, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (conn *MockConn) QueryRowsPartial(v any, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (conn *MockConn) QueryRowsPartialCtx(_ context.Context, v any, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (conn *MockConn) RawDB() (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func (conn *MockConn) Transact(func(session sqlx.Session) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (conn *MockConn) TransactCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (s statement) Close() error { _ = "STUB: not implemented"; return nil }

func (s statement) Exec(args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (s statement) ExecCtx(_ context.Context, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (s statement) QueryRow(v any, args ...any) error { _ = "STUB: not implemented"; return nil }

func (s statement) QueryRowCtx(_ context.Context, v any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s statement) QueryRowPartial(v any, args ...any) error { _ = "STUB: not implemented"; return nil }

func (s statement) QueryRowPartialCtx(_ context.Context, v any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s statement) QueryRows(v any, args ...any) error { _ = "STUB: not implemented"; return nil }

func (s statement) QueryRowsCtx(_ context.Context, v any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s statement) QueryRowsPartial(v any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s statement) QueryRowsPartialCtx(_ context.Context, v any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}
