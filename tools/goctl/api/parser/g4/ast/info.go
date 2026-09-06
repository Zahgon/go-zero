package ast

import "github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"

type InfoExpr struct {
	Info Expr
	Lp   Expr
	Rp   Expr
	Kvs  []*KvExpr
}

func (v *ApiVisitor) VisitInfoSpec(ctx *api.InfoSpecContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (i *InfoExpr) Format() error { _ = "STUB: not implemented"; return nil }

func (i *InfoExpr) Equal(v any) bool { _ = "STUB: not implemented"; return false }
