package ast

import "github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/token"

type InfoStmt struct {
	Info *TokenNode

	LParen *TokenNode

	Values []*KVExpr

	RParen *TokenNode
}

func (i *InfoStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (i *InfoStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (i *InfoStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (i *InfoStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (i *InfoStmt) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (i *InfoStmt) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (i *InfoStmt) stmtNode() { _ = "STUB: not implemented"; return }
