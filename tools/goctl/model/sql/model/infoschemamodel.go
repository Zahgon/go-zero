package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const indexPri = "PRIMARY"

type (
	InformationSchemaModel struct {
		conn sqlx.SqlConn
	}

	Column struct {
		*DbColumn
		Index *DbIndex
	}

	DbColumn struct {
		Name            string `db:"COLUMN_NAME"`
		DataType        string `db:"DATA_TYPE"`
		ColumnType      string `db:"COLUMN_TYPE"`
		Extra           string `db:"EXTRA"`
		Comment         string `db:"COLUMN_COMMENT"`
		ColumnDefault   any    `db:"COLUMN_DEFAULT"`
		IsNullAble      string `db:"IS_NULLABLE"`
		OrdinalPosition int    `db:"ORDINAL_POSITION"`
	}

	DbIndex struct {
		IndexName  string `db:"INDEX_NAME"`
		NonUnique  int    `db:"NON_UNIQUE"`
		SeqInIndex int    `db:"SEQ_IN_INDEX"`
	}

	ColumnData struct {
		Db      string
		Table   string
		Columns []*Column
	}

	Table struct {
		Db      string
		Table   string
		Columns []*Column

		UniqueIndex map[string][]*Column
		PrimaryKey  *Column
		NormalIndex map[string][]*Column
	}

	IndexType string

	Index struct {
		IndexType IndexType
		Columns   []*Column
	}
)

func NewInformationSchemaModel(conn sqlx.SqlConn) *InformationSchemaModel {
	_ = "STUB: not implemented"
	return nil
}

func (m *InformationSchemaModel) GetAllTables(database string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *InformationSchemaModel) FindColumns(db, table string) (*ColumnData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *InformationSchemaModel) FindIndex(db, table, column string) ([]*DbIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ColumnData) Convert() (*Table, error) { _ = "STUB: not implemented"; return nil, nil }
