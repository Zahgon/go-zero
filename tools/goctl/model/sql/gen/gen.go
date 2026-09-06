package gen

import (
	"bytes"

	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/model"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
)

const pwd = "."

type (
	defaultGenerator struct {
		console.Console

		dir           string
		pkg           string
		cfg           *config.Config
		isPostgreSql  bool
		ignoreColumns []string
		prefix        string
	}

	Option func(generator *defaultGenerator)

	code struct {
		importsCode    string
		varsCode       string
		typesCode      string
		newCode        string
		insertCode     string
		findCode       []string
		updateCode     string
		deleteCode     string
		cacheExtra     string
		tableName      string
		customizedCode string
	}

	codeTuple struct {
		modelCode       string
		modelCustomCode string
	}
)

func NewDefaultGenerator(prefix, dir string, cfg *config.Config, opt ...Option) (*defaultGenerator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithConsoleOption(c console.Console) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithIgnoreColumns(ignoreColumns []string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPostgreSql() Option { _ = "STUB: not implemented"; return *new(Option) }

func newDefaultOption() Option { _ = "STUB: not implemented"; return *new(Option) }

func (g *defaultGenerator) StartFromDDL(filename string, withCache, strict bool, database string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *defaultGenerator) StartFromInformationSchema(tables map[string]*model.Table, withCache, strict bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *defaultGenerator) createFile(modelList map[string]*codeTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *defaultGenerator) genFromDDL(filename string, withCache, strict bool, database string) (
	map[string]*codeTuple, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Table struct {
	parser.Table
	PrimaryCacheKey        Key
	UniqueCacheKey         []Key
	ContainsUniqueCacheKey bool
	ignoreColumns          []string
}

func (t Table) isIgnoreColumns(columnName string) bool { _ = "STUB: not implemented"; return false }

func (g *defaultGenerator) genModel(in parser.Table, withCache bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (g *defaultGenerator) genModelCustom(in parser.Table, withCache bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (g *defaultGenerator) executeModel(table Table, code *code) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wrapWithRawString(v string, postgreSql bool) string { _ = "STUB: not implemented"; return "" }
