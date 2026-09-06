package ast

import "github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"

type SyntaxExpr struct {
	Syntax      Expr
	Assign      Expr
	Version     Expr
	DocExpr     []Expr
	CommentExpr Expr
}

func (v *ApiVisitor) VisitSyntaxLit(ctx *api.SyntaxLitContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (s *SyntaxExpr) Format() error { _ = "STUB: not implemented"; return nil }

func (s *SyntaxExpr) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (s *SyntaxExpr) Doc() []Expr { _ = "STUB: not implemented"; return nil }

func (s *SyntaxExpr) Comment() Expr { _ = "STUB: not implemented"; return *new(Expr) }
