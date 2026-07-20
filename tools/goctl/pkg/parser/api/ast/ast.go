package ast

import (
	"io"

	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/token"
)

type Node interface {
	Pos() token.Position

	End() token.Position

	Format(...string) string

	HasHeadCommentGroup() bool

	HasLeadingCommentGroup() bool

	CommentGroup() (head, leading CommentGroup)
}

type Stmt interface {
	Node
	stmtNode()
}

type Expr interface {
	Node
	exprNode()
}

type AST struct {
	Filename     string
	Stmts        []Stmt
	readPosition int
}

type TokenNode struct {
	HeadCommentGroup CommentGroup

	Token token.Token

	LeadingCommentGroup CommentGroup

	headFlag, leadingFlag bool
}

func NewTokenNode(tok token.Token) *TokenNode { _ = "STUB: not implemented"; return nil }

func (t *TokenNode) IsEmptyString() bool { _ = "STUB: not implemented"; return false }

func (t *TokenNode) IsZeroString() bool { _ = "STUB: not implemented"; return false }

func (t *TokenNode) Equal(s string) bool { _ = "STUB: not implemented"; return false }

func (t *TokenNode) SetLeadingCommentGroup(cg CommentGroup) { _ = "STUB: not implemented"; return }

func (t *TokenNode) RawText() string { _ = "STUB: not implemented"; return "" }

func (t *TokenNode) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *TokenNode) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *TokenNode) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (t *TokenNode) PeekFirstLeadingComment() *CommentStmt { _ = "STUB: not implemented"; return nil }

func (t *TokenNode) PeekFirstHeadComment() *CommentStmt { _ = "STUB: not implemented"; return nil }

func (t *TokenNode) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (t *TokenNode) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (t *TokenNode) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (a *AST) Format(w io.Writer) { _ = "STUB: not implemented"; return }

func (a *AST) FormatForUnitTest(w io.Writer) { _ = "STUB: not implemented"; return }

func (a *AST) Print() { _ = "STUB: not implemented"; return }

func SyntaxError(pos token.Position, format string, v ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func DuplicateStmtError(pos token.Position, msg string) error {
	_ = "STUB: not implemented"
	return nil
}

func peekOne(list []string) string { _ = "STUB: not implemented"; return "" }
