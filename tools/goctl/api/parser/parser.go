package parser

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

type parser struct {
	ast  *ast.Api
	spec *spec.ApiSpec
}

func Parse(filename string) (*spec.ApiSpec, error) { _ = "STUB: not implemented"; return nil, nil }

func parseContent(content string, skipCheckTypeDeclaration bool, filename ...string) (*spec.ApiSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseContent(content string, filename ...string) (*spec.ApiSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseContentWithParserSkipCheckTypeDeclaration(content string, filename ...string) (*spec.ApiSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p parser) convert2Spec() error { _ = "STUB: not implemented"; return nil }

func (p parser) fillInfo() { _ = "STUB: not implemented"; return }

func (p parser) fillSyntax() { _ = "STUB: not implemented"; return }

func (p parser) fillImport() { _ = "STUB: not implemented"; return }

func (p parser) fillTypes() error { _ = "STUB: not implemented"; return nil }

func (p parser) findDefinedType(name string) (*spec.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p parser) fieldToMember(field *ast.TypeField) spec.Member {
	_ = "STUB: not implemented"
	return *new(spec.Member)
}

func (p parser) astTypeToSpec(in ast.DataType) spec.Type {
	_ = "STUB: not implemented"
	return *new(spec.Type)
}

func (p parser) stringExprs(docs []ast.Expr) []string { _ = "STUB: not implemented"; return nil }

func (p parser) commentExprs(comment ast.Expr) string { _ = "STUB: not implemented"; return "" }

func (p parser) fillService() error { _ = "STUB: not implemented"; return nil }

func (p parser) fillRouteAtServer(astRoute *ast.ServiceRoute, route *spec.Route) error {
	_ = "STUB: not implemented"
	return nil
}

func (p parser) fillAtServer(item *ast.Service, group *spec.Group) {
	_ = "STUB: not implemented"
	return
}

func (p parser) fillRouteType(route *spec.Route) error { _ = "STUB: not implemented"; return nil }
