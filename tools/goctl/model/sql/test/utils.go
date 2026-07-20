package mocksql

import (
	"database/sql"
)

var ErrNotFound = sql.ErrNoRows

func escape(input string) string { _ = "STUB: not implemented"; return "" }

func format(query string, args ...any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func logSqlError(stmt string, err error) { _ = "STUB: not implemented"; return }
