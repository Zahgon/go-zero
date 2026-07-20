package ast

import "github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/token"

type SyntaxStmt struct {
	Syntax *TokenNode

	Assign *TokenNode

	Value *TokenNode
}

func (s *SyntaxStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (s *SyntaxStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (s *SyntaxStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (s *SyntaxStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (s *SyntaxStmt) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (s *SyntaxStmt) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (s *SyntaxStmt) stmtNode() { _ = "STUB: not implemented"; return }
