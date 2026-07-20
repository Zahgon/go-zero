package ast

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"
)

const (
	prefixKey = "prefix"
	groupKey  = "group"
)

type Api struct {
	LinePrefix string
	Syntax     *SyntaxExpr
	Import     []*ImportExpr
	importM    map[string]PlaceHolder
	Info       *InfoExpr
	Type       []TypeExpr
	typeM      map[string]PlaceHolder
	Service    []*Service
	serviceM   map[string]PlaceHolder
	handlerM   map[string]PlaceHolder
	routeM     map[string]PlaceHolder
}

func (v *ApiVisitor) VisitApi(ctx *api.ApiContext) any { _ = "STUB: not implemented"; return *new(any) }

func (v *ApiVisitor) acceptService(root, final *Api) { _ = "STUB: not implemented"; return }

func (v *ApiVisitor) duplicateServerItemCheck(service *Service) { _ = "STUB: not implemented"; return }

func (v *ApiVisitor) acceptType(root, final *Api) { _ = "STUB: not implemented"; return }

func (v *ApiVisitor) acceptInfo(root, final *Api) { _ = "STUB: not implemented"; return }

func (v *ApiVisitor) acceptImport(root, final *Api) { _ = "STUB: not implemented"; return }

func (v *ApiVisitor) acceptSyntax(root, final *Api) { _ = "STUB: not implemented"; return }

func (v *ApiVisitor) VisitSpec(ctx *api.SpecContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (a *Api) Format() error { _ = "STUB: not implemented"; return nil }

func (a *Api) Equal(v any) bool { _ = "STUB: not implemented"; return false }
