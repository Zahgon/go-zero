package postgres

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const postgresDriverName = "pgx"

func New(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	_ = "STUB: not implemented"
	return *new(sqlx.SqlConn)
}
