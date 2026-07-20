package sqlx

import (
	"database/sql"
	"time"

	"github.com/zeromicro/go-zero/core/syncx"
)

const (
	maxIdleConns = 64
	maxOpenConns = 64
	maxLifetime  = time.Minute
)

var connManager = syncx.NewResourceManager()

func getCachedSqlConn(driverName, server string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSqlConn(driverName, server string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDBConnection(driverName, datasource string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
