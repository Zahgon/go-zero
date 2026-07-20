package dbtest

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func RunTest(t *testing.T, fn func(db *sql.DB, mock sqlmock.Sqlmock)) {
	_ = "STUB: not implemented"
	return
}

func RunTxTest(t *testing.T, f func(tx *sql.Tx, mock sqlmock.Sqlmock)) {
	_ = "STUB: not implemented"
	return
}
