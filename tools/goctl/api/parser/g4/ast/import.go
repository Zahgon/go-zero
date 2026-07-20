package ast

import "github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"

type ImportExpr struct {
	Import      Expr
	Value       Expr
	DocExpr     []Expr
	CommentExpr Expr
}

func (v *ApiVisitor) VisitImportSpec(ctx *api.ImportSpecContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitImportLit(ctx *api.ImportLitContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitImportBlock(ctx *api.ImportBlockContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitImportBlockValue(ctx *api.ImportBlockValueContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitImportValue(ctx *api.ImportValueContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (i *ImportExpr) Format() error { _ = "STUB: not implemented"; return nil }

func (i *ImportExpr) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (i *ImportExpr) Doc() []Expr { _ = "STUB: not implemented"; return nil }

func (i *ImportExpr) Comment() Expr { _ = "STUB: not implemented"; return *new(Expr) }
