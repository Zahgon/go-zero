package parser

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/ast"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/placeholder"
)

type Analyzer struct {
	api  *API
	spec *spec.ApiSpec
}

func (a *Analyzer) astTypeToSpec(in ast.DataType) (spec.Type, error) {
	_ = "STUB: not implemented"
	return *new(spec.Type), nil
}

func (a *Analyzer) convert2Spec() error { _ = "STUB: not implemented"; return nil }

func (a *Analyzer) convertAtDoc(atDoc ast.AtDocStmt) spec.AtDoc {
	_ = "STUB: not implemented"
	return *new(spec.AtDoc)
}

func (a *Analyzer) convertKV(kv []*ast.KVExpr) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (a *Analyzer) fieldToMember(field *ast.ElemExpr) (spec.Member, error) {
	_ = "STUB: not implemented"
	return *new(spec.Member), nil
}

func (a *Analyzer) fillRouteType(route *spec.Route) error { _ = "STUB: not implemented"; return nil }

func (a *Analyzer) fillService() error { _ = "STUB: not implemented"; return nil }

func (a *Analyzer) fillInfo() { _ = "STUB: not implemented"; return }

func (a *Analyzer) fillTypes() error { _ = "STUB: not implemented"; return nil }

func (a *Analyzer) fillTypeExpr(expr *ast.TypeExpr) error { _ = "STUB: not implemented"; return nil }

func (a *Analyzer) findDefinedType(name string) (spec.Type, error) {
	_ = "STUB: not implemented"
	return *new(spec.Type), nil
}

func (a *Analyzer) getType(expr *ast.BodyStmt, req bool) (spec.Type, error) {
	_ = "STUB: not implemented"
	return *new(spec.Type), nil
}

func Parse(filename string, src interface{}) (*spec.ApiSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var kind = map[string]placeholder.Type{
	"bool":       placeholder.PlaceHolder,
	"int":        placeholder.PlaceHolder,
	"int8":       placeholder.PlaceHolder,
	"int16":      placeholder.PlaceHolder,
	"int32":      placeholder.PlaceHolder,
	"int64":      placeholder.PlaceHolder,
	"uint":       placeholder.PlaceHolder,
	"uint8":      placeholder.PlaceHolder,
	"uint16":     placeholder.PlaceHolder,
	"uint32":     placeholder.PlaceHolder,
	"uint64":     placeholder.PlaceHolder,
	"uintptr":    placeholder.PlaceHolder,
	"float32":    placeholder.PlaceHolder,
	"float64":    placeholder.PlaceHolder,
	"complex64":  placeholder.PlaceHolder,
	"complex128": placeholder.PlaceHolder,
	"string":     placeholder.PlaceHolder,
	"byte":       placeholder.PlaceHolder,
	"rune":       placeholder.PlaceHolder,
	"any":        placeholder.PlaceHolder,
}

func IsBaseType(text string) bool { _ = "STUB: not implemented"; return false }
