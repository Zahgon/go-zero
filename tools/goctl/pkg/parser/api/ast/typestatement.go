package ast

import "github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/token"

type TypeStmt interface {
	Stmt
	typeNode()
}

type TypeLiteralStmt struct {
	Type *TokenNode

	Expr *TypeExpr
}

func (t *TypeLiteralStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *TypeLiteralStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *TypeLiteralStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (t *TypeLiteralStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (t *TypeLiteralStmt) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *TypeLiteralStmt) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *TypeLiteralStmt) stmtNode() { _ = "STUB: not implemented"; return }
func (t *TypeLiteralStmt) typeNode() { _ = "STUB: not implemented"; return }

type TypeGroupStmt struct {
	Type *TokenNode

	LParen *TokenNode

	ExprList []*TypeExpr

	RParen *TokenNode
}

func (t *TypeGroupStmt) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *TypeGroupStmt) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *TypeGroupStmt) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (t *TypeGroupStmt) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (t *TypeGroupStmt) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *TypeGroupStmt) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *TypeGroupStmt) stmtNode() { _ = "STUB: not implemented"; return }
func (t *TypeGroupStmt) typeNode() { _ = "STUB: not implemented"; return }

type TypeExpr struct {
	Name *TokenNode

	Assign *TokenNode

	DataType DataType
}

func (e *TypeExpr) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (e *TypeExpr) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (e *TypeExpr) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (e *TypeExpr) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (e *TypeExpr) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (e *TypeExpr) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (e *TypeExpr) exprNode() { _ = "STUB: not implemented"; return }

func (e *TypeExpr) isStruct() bool { _ = "STUB: not implemented"; return false }

type ElemExpr struct {
	Name []*TokenNode

	DataType DataType

	Tag *TokenNode
}

func (e *ElemExpr) IsAnonymous() bool { _ = "STUB: not implemented"; return false }

func (e *ElemExpr) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (e *ElemExpr) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (e *ElemExpr) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (e *ElemExpr) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (e *ElemExpr) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (e *ElemExpr) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (e *ElemExpr) exprNode() { _ = "STUB: not implemented"; return }

type ElemExprList []*ElemExpr

type DataType interface {
	Expr
	dataTypeNode()

	CanEqual() bool

	ContainsStruct() bool

	RawText() string
}

type AnyDataType struct {
	Any     *TokenNode
	isChild bool
}

func (t *AnyDataType) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *AnyDataType) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *AnyDataType) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (t *AnyDataType) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (t *AnyDataType) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (t *AnyDataType) RawText() string { _ = "STUB: not implemented"; return "" }

func (t *AnyDataType) ContainsStruct() bool { _ = "STUB: not implemented"; return false }

func (t *AnyDataType) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (t *AnyDataType) exprNode() { _ = "STUB: not implemented"; return }

func (t *AnyDataType) dataTypeNode() { _ = "STUB: not implemented"; return }

func (t *AnyDataType) CanEqual() bool { _ = "STUB: not implemented"; return false }

type ArrayDataType struct {
	LBrack *TokenNode

	Length *TokenNode

	RBrack *TokenNode

	DataType DataType
	isChild  bool
}

func (t *ArrayDataType) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *ArrayDataType) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *ArrayDataType) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (t *ArrayDataType) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (t *ArrayDataType) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *ArrayDataType) RawText() string { _ = "STUB: not implemented"; return "" }

func (t *ArrayDataType) ContainsStruct() bool { _ = "STUB: not implemented"; return false }

func (t *ArrayDataType) CanEqual() bool { _ = "STUB: not implemented"; return false }

func (t *ArrayDataType) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *ArrayDataType) exprNode()     { _ = "STUB: not implemented"; return }
func (t *ArrayDataType) dataTypeNode() { _ = "STUB: not implemented"; return }

type BaseDataType struct {
	Base    *TokenNode
	isChild bool
}

func (t *BaseDataType) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *BaseDataType) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *BaseDataType) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (t *BaseDataType) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (t *BaseDataType) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (t *BaseDataType) RawText() string { _ = "STUB: not implemented"; return "" }

func (t *BaseDataType) ContainsStruct() bool { _ = "STUB: not implemented"; return false }

func (t *BaseDataType) CanEqual() bool { _ = "STUB: not implemented"; return false }

func (t *BaseDataType) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (t *BaseDataType) exprNode()     { _ = "STUB: not implemented"; return }
func (t *BaseDataType) dataTypeNode() { _ = "STUB: not implemented"; return }

type InterfaceDataType struct {
	Interface *TokenNode
	isChild   bool
}

func (t *InterfaceDataType) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *InterfaceDataType) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *InterfaceDataType) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (t *InterfaceDataType) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (t *InterfaceDataType) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *InterfaceDataType) RawText() string { _ = "STUB: not implemented"; return "" }

func (t *InterfaceDataType) ContainsStruct() bool { _ = "STUB: not implemented"; return false }

func (t *InterfaceDataType) CanEqual() bool { _ = "STUB: not implemented"; return false }

func (t *InterfaceDataType) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *InterfaceDataType) exprNode() { _ = "STUB: not implemented"; return }

func (t *InterfaceDataType) dataTypeNode() { _ = "STUB: not implemented"; return }

type MapDataType struct {
	Map *TokenNode

	LBrack *TokenNode

	Key DataType

	RBrack *TokenNode

	Value   DataType
	isChild bool
}

func (t *MapDataType) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *MapDataType) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *MapDataType) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (t *MapDataType) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (t *MapDataType) End() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (t *MapDataType) RawText() string { _ = "STUB: not implemented"; return "" }

func (t *MapDataType) ContainsStruct() bool { _ = "STUB: not implemented"; return false }

func (t *MapDataType) CanEqual() bool { _ = "STUB: not implemented"; return false }

func (t *MapDataType) Pos() token.Position { _ = "STUB: not implemented"; return *new(token.Position) }

func (t *MapDataType) exprNode()     { _ = "STUB: not implemented"; return }
func (t *MapDataType) dataTypeNode() { _ = "STUB: not implemented"; return }

type PointerDataType struct {
	Star *TokenNode

	DataType DataType
	isChild  bool
}

func (t *PointerDataType) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *PointerDataType) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *PointerDataType) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (t *PointerDataType) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (t *PointerDataType) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *PointerDataType) RawText() string { _ = "STUB: not implemented"; return "" }

func (t *PointerDataType) ContainsStruct() bool { _ = "STUB: not implemented"; return false }

func (t *PointerDataType) CanEqual() bool { _ = "STUB: not implemented"; return false }

func (t *PointerDataType) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *PointerDataType) exprNode()     { _ = "STUB: not implemented"; return }
func (t *PointerDataType) dataTypeNode() { _ = "STUB: not implemented"; return }

type SliceDataType struct {
	LBrack *TokenNode

	RBrack *TokenNode

	DataType DataType
	isChild  bool
}

func (t *SliceDataType) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *SliceDataType) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *SliceDataType) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (t *SliceDataType) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (t *SliceDataType) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *SliceDataType) RawText() string { _ = "STUB: not implemented"; return "" }

func (t *SliceDataType) ContainsStruct() bool { _ = "STUB: not implemented"; return false }

func (t *SliceDataType) CanEqual() bool { _ = "STUB: not implemented"; return false }

func (t *SliceDataType) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *SliceDataType) exprNode()     { _ = "STUB: not implemented"; return }
func (t *SliceDataType) dataTypeNode() { _ = "STUB: not implemented"; return }

type StructDataType struct {
	LBrace *TokenNode

	Elements ElemExprList

	RBrace  *TokenNode
	isChild bool
}

func (t *StructDataType) HasHeadCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *StructDataType) HasLeadingCommentGroup() bool { _ = "STUB: not implemented"; return false }

func (t *StructDataType) CommentGroup() (head, leading CommentGroup) {
	_ = "STUB: not implemented"
	return *new(CommentGroup), *new(CommentGroup)
}

func (t *StructDataType) Format(prefix ...string) string { _ = "STUB: not implemented"; return "" }

func (t *StructDataType) End() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *StructDataType) RawText() string { _ = "STUB: not implemented"; return "" }

func (t *StructDataType) ContainsStruct() bool { _ = "STUB: not implemented"; return false }

func (t *StructDataType) CanEqual() bool { _ = "STUB: not implemented"; return false }

func (t *StructDataType) Pos() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (t *StructDataType) exprNode()     { _ = "STUB: not implemented"; return }
func (t *StructDataType) dataTypeNode() { _ = "STUB: not implemented"; return }
