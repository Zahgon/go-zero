package ast

import "github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/token"

type ImportStmt interface {
	Stmt
	importNode()
}

type ImportLiteralStmt struct {
	Import *TokenNode

	Value *TokenNode
}

func (i *ImportLiteralStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (i *ImportLiteralStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (i *ImportLiteralStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (i *ImportLiteralStmt) Format(prefix ...string) (result string) {
	_ = "STUB: not implemented"
	return ""
}

func (i *ImportLiteralStmt) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (i *ImportLiteralStmt) importNode() { _ = "STUB: not implemented"; return }

func (i *ImportLiteralStmt) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (i *ImportLiteralStmt) stmtNode() { _ = "STUB: not implemented"; return }

type ImportGroupStmt struct {
	Import *TokenNode

	LParen *TokenNode

	Values []*TokenNode

	RParen *TokenNode
}

func (i *ImportGroupStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (i *ImportGroupStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (i *ImportGroupStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (i *ImportGroupStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (i *ImportGroupStmt) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (i *ImportGroupStmt) importNode() { _ = "STUB: not implemented"; return }

func (i *ImportGroupStmt) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (i *ImportGroupStmt) stmtNode() { _ = "STUB: not implemented"; return }
