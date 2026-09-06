package ast

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"
)

type Service struct {
	AtServer   *AtServer
	ServiceApi *ServiceApi
}

type KV []*KvExpr

type AtServer struct {
	AtServerToken Expr
	Lp            Expr
	Rp            Expr
	Kv            KV
}

type ServiceApi struct {
	ServiceToken Expr
	Name         Expr
	Lbrace       Expr
	Rbrace       Expr
	ServiceRoute []*ServiceRoute
}

type ServiceRoute struct {
	AtDoc     *AtDoc
	AtServer  *AtServer
	AtHandler *AtHandler
	Route     *Route
}

type AtDoc struct {
	AtDocToken Expr
	Lp         Expr
	Rp         Expr
	LineDoc    Expr
	Kv         []*KvExpr
}

type AtHandler struct {
	AtHandlerToken Expr
	Name           Expr
	DocExpr        []Expr
	CommentExpr    Expr
}

type Route struct {
	Method      Expr
	Path        Expr
	Req         *Body
	ReturnToken Expr
	Reply       *Body
	DocExpr     []Expr
	CommentExpr Expr
}

type Body struct {
	ReturnExpr Expr
	Lp         Expr
	Rp         Expr
	Name       DataType
}

func (v *ApiVisitor) VisitServiceSpec(ctx *api.ServiceSpecContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitAtServer(ctx *api.AtServerContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitServiceApi(ctx *api.ServiceApiContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitServiceRoute(ctx *api.ServiceRouteContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitAtDoc(ctx *api.AtDocContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitAtHandler(ctx *api.AtHandlerContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitRoute(ctx *api.RouteContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitBody(ctx *api.BodyContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitReplybody(ctx *api.ReplybodyContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (b *Body) Format() error { _ = "STUB: not implemented"; return nil }

func (b *Body) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (r *Route) Format() error { _ = "STUB: not implemented"; return nil }

func (r *Route) Doc() []Expr { _ = "STUB: not implemented"; return nil }

func (r *Route) Comment() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (r *Route) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (a *AtHandler) Doc() []Expr { _ = "STUB: not implemented"; return nil }

func (a *AtHandler) Comment() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (a *AtHandler) Format() error { _ = "STUB: not implemented"; return nil }

func (a *AtHandler) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (a *AtDoc) Format() error { _ = "STUB: not implemented"; return nil }

func (a *AtDoc) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (a *AtServer) Format() error { _ = "STUB: not implemented"; return nil }

func (a *AtServer) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (s *ServiceRoute) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (s *ServiceRoute) Format() error { _ = "STUB: not implemented"; return nil }

func (s *ServiceRoute) GetHandler() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (a *ServiceApi) Format() error { _ = "STUB: not implemented"; return nil }

func (a *ServiceApi) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (s *Service) Format() error { _ = "STUB: not implemented"; return nil }

func (s *Service) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (kv KV) Get(key string) Expr { _ = "STUB: not implemented"; return *new(Expr) }
