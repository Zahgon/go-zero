package ast

import (
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/token"
)

type CommentGroup []*CommentStmt

func (cg CommentGroup) List() []string { _ = "STUB: not implemented"; return nil }

func (cg CommentGroup) String() string { _ = "STUB: not implemented"; return "" }

func (cg CommentGroup) Join(sep string) string { _ = "STUB: not implemented"; return "" }

func (cg CommentGroup) Valid() bool { _ = "STUB: not implemented"; return false }

type CommentStmt struct {
	Comment token.Token
}

func (c *CommentStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (c *CommentStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (c *CommentStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (c *CommentStmt) stmtNode() { _ = "STUB: not implemented"; return }

func (c *CommentStmt) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (c *CommentStmt) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (c *CommentStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }
