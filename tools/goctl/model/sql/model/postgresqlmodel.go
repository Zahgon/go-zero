package model

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var p2m = map[string]string{
	"int8":        "bigint",
	"float8":      "double",
	"float4":      "float",
	"int2":        "smallint",
	"int4":        "integer",
	"timestamptz": "timestamp",
	"uuid":        "varchar",
}

type PostgreSqlModel struct {
	conn sqlx.SqlConn
}

type PostgreColumn struct {
	Num               sql.NullInt32  `db:"num"`
	Field             sql.NullString `db:"field"`
	Type              sql.NullString `db:"type"`
	NotNull           sql.NullBool   `db:"not_null"`
	Comment           sql.NullString `db:"comment"`
	ColumnDefault     sql.NullString `db:"column_default"`
	IdentityIncrement sql.NullInt32  `db:"identity_increment"`
}

type PostgreIndex struct {
	IndexName  sql.NullString `db:"index_name"`
	IndexId    sql.NullInt32  `db:"index_id"`
	IsUnique   sql.NullBool   `db:"is_unique"`
	IsPrimary  sql.NullBool   `db:"is_primary"`
	ColumnName sql.NullString `db:"column_name"`
	IndexSort  sql.NullInt32  `db:"index_sort"`
}

func NewPostgreSqlModel(conn sqlx.SqlConn) *PostgreSqlModel { _ = "STUB: not implemented"; return nil }

func (m *PostgreSqlModel) GetAllTables(schema string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *PostgreSqlModel) FindColumns(schema, table string) (*ColumnData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *PostgreSqlModel) getColumns(schema, table string, in []*PostgreColumn) ([]*Column, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *PostgreSqlModel) convertPostgreSqlTypeIntoMysqlType(in string) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *PostgreSqlModel) getIndex(schema, table string) (map[string][]*DbIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *PostgreSqlModel) FindIndex(schema, table string) ([]*PostgreIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
