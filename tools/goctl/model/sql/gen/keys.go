package gen

import (
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
)

type Key struct {
	VarLeft string

	VarRight string

	VarExpression string

	KeyLeft string

	KeyRight string

	DataKeyRight string

	KeyExpression string

	DataKeyExpression string

	FieldNameJoin Join

	Fields []*parser.Field
}

type Join []string

func genCacheKeys(prefix string, table parser.Table) (Key, []Key) {
	_ = "STUB: not implemented"
	return *new(Key), nil
}

func genCacheKey(prefix string, db, table stringx.String, in []*parser.Field) Key {
	_ = "STUB: not implemented"
	return *new(Key)
}

func (j Join) Title() Join { _ = "STUB: not implemented"; return *new(Join) }

func (j Join) Camel() Join { _ = "STUB: not implemented"; return *new(Join) }

func (j Join) Snake() Join { _ = "STUB: not implemented"; return *new(Join) }

func (j Join) Untitle() Join { _ = "STUB: not implemented"; return *new(Join) }

func (j Join) Upper() Join { _ = "STUB: not implemented"; return *new(Join) }

func (j Join) Lower() Join { _ = "STUB: not implemented"; return *new(Join) }

func (j Join) With(sep string) stringx.String {
	_ = "STUB: not implemented"
	return *new(stringx.String)
}
