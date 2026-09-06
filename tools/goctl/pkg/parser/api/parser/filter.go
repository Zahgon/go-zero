package parser

import (
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/ast"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/token"
)

type filterBuilder struct {
	filename      string
	m             map[string]token.Position
	checkExprName string
	errorManager  *errorManager
}

func (b *filterBuilder) check(nodes ...*ast.TokenNode) { _ = "STUB: not implemented"; return }

func (b *filterBuilder) checkNodeWithPrefix(prefix string, nodes ...*ast.TokenNode) {
	_ = "STUB: not implemented"
	return
}

func (b *filterBuilder) error() error { _ = "STUB: not implemented"; return nil }

type filter struct {
	builders []*filterBuilder
}

func newFilter() *filter { _ = "STUB: not implemented"; return nil }

func (f *filter) addCheckItem(filename, checkExprName string) *filterBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (f *filter) error() error { _ = "STUB: not implemented"; return nil }
