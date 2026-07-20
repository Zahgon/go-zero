package mocksql

import (
	"database/sql"
	"time"
)

const slowThreshold = time.Millisecond * 500

func exec(db *sql.DB, q string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func execStmt(conn *sql.Stmt, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func query(db *sql.DB, scanner func(*sql.Rows) error, q string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func queryStmt(conn *sql.Stmt, scanner func(*sql.Rows) error, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}
