package ast

import "github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/token"

type AtServerStmt struct {
	AtServer *TokenNode

	LParen *TokenNode

	Values []*KVExpr

	RParen *TokenNode
}

func (a *AtServerStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (a *AtServerStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (a *AtServerStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (a *AtServerStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (a *AtServerStmt) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (a *AtServerStmt) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (a *AtServerStmt) stmtNode() { _ = "STUB: not implemented"; return }

type AtDocStmt interface {
	Stmt
	atDocNode()
}

type AtDocLiteralStmt struct {
	AtDoc *TokenNode
	Value *TokenNode
}

func (a *AtDocLiteralStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (a *AtDocLiteralStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (a *AtDocLiteralStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (a *AtDocLiteralStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (a *AtDocLiteralStmt) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (a *AtDocLiteralStmt) atDocNode() { _ = "STUB: not implemented"; return }

func (a *AtDocLiteralStmt) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (a *AtDocLiteralStmt) stmtNode() { _ = "STUB: not implemented"; return }

type AtDocGroupStmt struct {
	AtDoc  *TokenNode
	LParen *TokenNode
	Values []*KVExpr
	RParen *TokenNode
}

func (a *AtDocGroupStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (a *AtDocGroupStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (a *AtDocGroupStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (a *AtDocGroupStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (a *AtDocGroupStmt) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (a *AtDocGroupStmt) atDocNode() { _ = "STUB: not implemented"; return }

func (a *AtDocGroupStmt) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (a *AtDocGroupStmt) stmtNode() { _ = "STUB: not implemented"; return }

type ServiceStmt struct {
	AtServerStmt *AtServerStmt
	Service      *TokenNode
	Name         *ServiceNameExpr
	LBrace       *TokenNode
	Routes       []*ServiceItemStmt
	RBrace       *TokenNode
}

func (s *ServiceStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (s *ServiceStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (s *ServiceStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (s *ServiceStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (s *ServiceStmt) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (s *ServiceStmt) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (s *ServiceStmt) stmtNode() { _ = "STUB: not implemented"; return }

type ServiceNameExpr struct {
	Name *TokenNode
}

func (s *ServiceNameExpr) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (s *ServiceNameExpr) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (s *ServiceNameExpr) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (s *ServiceNameExpr) Format(...string) string { _ = "STUB: not implemented"; return "" }

func (s *ServiceNameExpr) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (s *ServiceNameExpr) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (s *ServiceNameExpr) exprNode() { _ = "STUB: not implemented"; return }

type AtHandlerStmt struct {
	AtHandler *TokenNode
	Name      *TokenNode
}

func (a *AtHandlerStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (a *AtHandlerStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (a *AtHandlerStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (a *AtHandlerStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (a *AtHandlerStmt) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (a *AtHandlerStmt) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (a *AtHandlerStmt) stmtNode() { _ = "STUB: not implemented"; return }

type ServiceItemStmt struct {
	AtDoc     AtDocStmt
	AtHandler *AtHandlerStmt
	Route     *RouteStmt
}

func (s *ServiceItemStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (s *ServiceItemStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (s *ServiceItemStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (s *ServiceItemStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (s *ServiceItemStmt) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (s *ServiceItemStmt) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (s *ServiceItemStmt) stmtNode() { _ = "STUB: not implemented"; return }

type RouteStmt struct {
	Method   *TokenNode
	Path     *PathExpr
	Request  *BodyStmt
	Returns  *TokenNode
	Response *BodyStmt
}

func (r *RouteStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (r *RouteStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (r *RouteStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (r *RouteStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (r *RouteStmt) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (r *RouteStmt) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (r *RouteStmt) stmtNode() { _ = "STUB: not implemented"; return }

type PathExpr struct {
	Value *TokenNode
}

func (p *PathExpr) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (p *PathExpr) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (p *PathExpr) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (p *PathExpr) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (p *PathExpr) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (p *PathExpr) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (p *PathExpr) exprNode() { _ = "STUB: not implemented"; return }

type BodyStmt struct {
	LParen *TokenNode
	Body   *BodyExpr
	RParen *TokenNode
}

func (b *BodyStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (b *BodyStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (b *BodyStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (b *BodyStmt) Format(...string) string { _ = "STUB: not implemented"; return "" }

func (b *BodyStmt) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (b *BodyStmt) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (b *BodyStmt) stmtNode() { _ = "STUB: not implemented"; return }

type BodyExpr struct {
	LBrack *TokenNode
	RBrack *TokenNode
	Star   *TokenNode
	Value  *TokenNode
}

func (e *BodyExpr) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (e *BodyExpr) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (e *BodyExpr) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (e *BodyExpr) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (e *BodyExpr) Format(...string) string { _ = "STUB: not implemented"; return "" }

func (e *BodyExpr) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (e *BodyExpr) exprNode() { _ = "STUB: not implemented"; return }

func (e *BodyExpr) IsArrayType() bool { _ = "STUB: not implemented"; return false }
