package ast

import "github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/token"

type KVExpr struct {
	Key *TokenNode

	Colon *TokenNode

	Value *TokenNode
}

func (i *KVExpr) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (i *KVExpr) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (i *KVExpr) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (i *KVExpr) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (i *KVExpr) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (i *KVExpr) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (i *KVExpr) exprNode() { _ = "STUB: not implemented"; return }
