package sqlx

const (
	mysqlDriverName           = "mysql"
	duplicateEntryCode uint16 = 1062
)

func NewMysql(datasource string, opts ...SqlOption) SqlConn {
	_ = "STUB: not implemented"
	return *new(SqlConn)
}

func mysqlAcceptable(err error) bool { _ = "STUB: not implemented"; return false }

func withMysqlAcceptable() SqlOption { _ = "STUB: not implemented"; return *new(SqlOption) }
