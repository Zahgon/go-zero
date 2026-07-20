package ast

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"
)

type KvExpr struct {
	Key         Expr
	Value       Expr
	DocExpr     []Expr
	CommentExpr Expr
}

func (v *ApiVisitor) VisitKvLit(ctx *api.KvLitContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (k *KvExpr) Format() error { _ = "STUB: not implemented"; return nil }

func (k *KvExpr) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (k *KvExpr) Doc() []Expr { _ = "STUB: not implemented"; return nil }

func (k *KvExpr) Comment() Expr { _ = "STUB: not implemented"; return *new(Expr) }
