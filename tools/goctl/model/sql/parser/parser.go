package parser

import (
	"github.com/zeromicro/ddl-parser/parser"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/model"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
)

const timeImport = "time.Time"

type (
	Table struct {
		Name        stringx.String
		Db          stringx.String
		PrimaryKey  Primary
		UniqueIndex map[string][]*Field
		Fields      []*Field
		ContainsPQ  bool
	}

	Primary struct {
		Field
		AutoIncrement bool
	}

	Field struct {
		NameOriginal    string
		Name            stringx.String
		ThirdPkg        string
		DataType        string
		Comment         string
		SeqInIndex      int
		OrdinalPosition int
		ContainsPQ      bool
	}

	KeyType int
)

func parseNameOriginal(ts []*parser.Table) (nameOriginals [][]string) {
	_ = "STUB: not implemented"
	return nil
}

func Parse(filename, database string, strict bool) ([]*Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkDuplicateUniqueIndex(uniqueIndex map[string][]*Field, tableName string) {
	_ = "STUB: not implemented"
	return
}

func convertColumns(columns []*parser.Column, primaryColumn string, strict bool) (Primary, map[string]*Field, error) {
	_ = "STUB: not implemented"
	return *new(Primary), nil, nil
}

func (t *Table) ContainsTime() bool { _ = "STUB: not implemented"; return false }

func ConvertDataType(table *model.Table, strict bool) (*Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTableFields(table *model.Table, strict bool) (map[string]*Field, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
