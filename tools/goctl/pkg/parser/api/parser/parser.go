package parser

import (
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/ast"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/scanner"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/token"
)

const (
	idAPI          = "api"
	groupKeyText   = "group"
	infoTitleKey   = "Title"
	infoDescKey    = "Desc"
	infoVersionKey = "Version"
	infoAuthorKey  = "Author"
	infoEmailKey   = "Email"
)

type Parser struct {
	s      *scanner.Scanner
	errors []error

	curTok  token.Token
	peekTok token.Token

	headCommentGroup ast.CommentGroup
	api              *ast.AST
	node             map[token.Token]*ast.TokenNode
}

func New(filename string, src interface{}) *Parser { _ = "STUB: not implemented"; return nil }

func (p *Parser) Parse() *ast.AST { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseStmt() ast.Stmt { _ = "STUB: not implemented"; return *new(ast.Stmt) }

func (p *Parser) parseService() *ast.ServiceStmt { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseServiceItemsStmt() []*ast.ServiceItemStmt {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) parseServiceItemStmt() *ast.ServiceItemStmt { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseRouteStmt() *ast.RouteStmt { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseBodyStmt() *ast.BodyStmt { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseBodyExpr() *ast.BodyExpr { _ = "STUB: not implemented"; return nil }

func (p *Parser) parsePathExpr() *ast.PathExpr { _ = "STUB: not implemented"; return nil }

func (p *Parser) parsePathItem() []token.Token { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseServiceNameExpr() *ast.ServiceNameExpr { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseAtDocStmt() ast.AtDocStmt {
	_ = "STUB: not implemented"
	return *new(ast.AtDocStmt)
}

func (p *Parser) parseAtDocGroupStmt() ast.AtDocStmt {
	_ = "STUB: not implemented"
	return *new(ast.AtDocStmt)
}

func (p *Parser) parseAtDocLiteralStmt() ast.AtDocStmt {
	_ = "STUB: not implemented"
	return *new(ast.AtDocStmt)
}

func (p *Parser) parseAtHandlerStmt() *ast.AtHandlerStmt { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseAtServerStmt() *ast.AtServerStmt { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseTypeStmt() ast.TypeStmt { _ = "STUB: not implemented"; return *new(ast.TypeStmt) }

func (p *Parser) parseTypeLiteralStmt() ast.TypeStmt {
	_ = "STUB: not implemented"
	return *new(ast.TypeStmt)
}

func (p *Parser) parseTypeGroupStmt() ast.TypeStmt {
	_ = "STUB: not implemented"
	return *new(ast.TypeStmt)
}

func (p *Parser) parseTypeExprList() []*ast.TypeExpr { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseTypeExpr() *ast.TypeExpr { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseDataType() ast.DataType { _ = "STUB: not implemented"; return *new(ast.DataType) }

func (p *Parser) parseStructDataType() *ast.StructDataType { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseElemExprList() ast.ElemExprList {
	_ = "STUB: not implemented"
	return *new(ast.ElemExprList)
}

func (p *Parser) parseElemExpr() *ast.ElemExpr { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseAnyDataType() *ast.AnyDataType { _ = "STUB: not implemented"; return nil }

func (p *Parser) parsePointerDataType() *ast.PointerDataType { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseInterfaceDataType() *ast.InterfaceDataType {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) parseMapDataType() *ast.MapDataType { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseArrayDataType() *ast.ArrayDataType { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseSliceDataType() *ast.SliceDataType { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseImportStmt() ast.ImportStmt {
	_ = "STUB: not implemented"
	return *new(ast.ImportStmt)
}

func (p *Parser) parseImportLiteralStmt() ast.ImportStmt {
	_ = "STUB: not implemented"
	return *new(ast.ImportStmt)
}

func (p *Parser) parseImportGroupStmt() ast.ImportStmt {
	_ = "STUB: not implemented"
	return *new(ast.ImportStmt)
}

func (p *Parser) parseInfoStmt() *ast.InfoStmt { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseAtServerKVExpression() *ast.KVExpr { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseKVExpression() *ast.KVExpr { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseSyntaxStmt() *ast.SyntaxStmt { _ = "STUB: not implemented"; return nil }

func (p *Parser) curTokenIsNotEof() bool { _ = "STUB: not implemented"; return false }

func (p *Parser) curTokenIsNot(expected token.Type) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) curTokenIsKeyword() bool { _ = "STUB: not implemented"; return false }

func (p *Parser) curTokenIs(expected ...interface{}) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) advanceIfPeekTokenIs(expected ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Parser) peekTokenIs(expected ...interface{}) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) peekTokenIsNot(expected ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Parser) notExpectPeekToken(expected ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Parser) notExpectPeekTokenGotComment(actual *ast.CommentStmt, expected ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Parser) expectPeekToken(expected ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Parser) expectIdentError(tok token.Token, expected ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (p *Parser) init() bool {
	if !p.nextToken() {
		return false
	}

	return p.nextToken()
}

func (p *Parser) nextToken() bool { _ = "STUB: not implemented"; return false }

func (p *Parser) curTokenNode() *ast.TokenNode { _ = "STUB: not implemented"; return nil }

func (p *Parser) getNode(tok token.Token) *ast.TokenNode { _ = "STUB: not implemented"; return nil }

func isNil(v interface{}) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) CheckErrors() error { _ = "STUB: not implemented"; return nil }

func (p *Parser) appendStmt(stmt ...ast.Stmt) { _ = "STUB: not implemented"; return }

func (p *Parser) hasNoErrors() bool { _ = "STUB: not implemented"; return false }

func (p *Parser) ParseForUintTest() *ast.AST { _ = "STUB: not implemented"; return nil }

func (p *Parser) parseStmtForUniTest() ast.Stmt { _ = "STUB: not implemented"; return *new(ast.Stmt) }
