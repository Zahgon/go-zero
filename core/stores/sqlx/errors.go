package sqlx

import (
	"database/sql"
	"errors"
)

var (
	ErrNotFound = sql.ErrNoRows

	errCantNestTx    = errors.New("cannot nest transactions")
	errNoRawDBFromTx = errors.New("cannot get raw db from transaction")
)
