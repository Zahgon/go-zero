package ast

import (
	"github.com/zeromicro/antlr"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
)

type (
	Parser struct {
		antlr.DefaultErrorListener
		linePrefix               string
		debug                    bool
		log                      console.Console
		skipCheckTypeDeclaration bool
		handlerMap               map[string]PlaceHolder
		routeMap                 map[string]PlaceHolder
		typeMap                  map[string]PlaceHolder
		fileMap                  map[string]PlaceHolder
		importStatck             importStack
		syntax                   *SyntaxExpr
	}

	ParserOption func(p *Parser)
)

func NewParser(options ...ParserOption) *Parser { _ = "STUB: not implemented"; return nil }

func (p *Parser) Accept(fn func(p *api.ApiParserParser, visitor *ApiVisitor) any, content string) (v any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (p *Parser) Parse(filename string) (*Api, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Parser) ParseContent(content string, filename ...string) (*Api, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) parse(filename, content string) (*Api, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) invokeImportedApi(filename string, imports []*ImportExpr) ([]*Api, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) alreadyImported(filename string) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) invoke(linePrefix, content string) (v *Api, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) storeVerificationInfo(api *Api) { _ = "STUB: not implemented"; return }

func (p *Parser) valid(nestedApi *Api) error { _ = "STUB: not implemented"; return nil }

func (p *Parser) duplicateRouteCheck(nestedApi *Api) error { _ = "STUB: not implemented"; return nil }

func (p *Parser) nestedApiCheck(mainApi, nestedApi *Api) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) memberFill(apiList []*Api) *Api { _ = "STUB: not implemented"; return nil }

func (p *Parser) checkTypeDeclaration(apiList []*Api) error { _ = "STUB: not implemented"; return nil }

func (p *Parser) checkServices(apiItem *Api, types map[string]TypeExpr, linePrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) checkRequestBody(route *Route, types map[string]TypeExpr, linePrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) checkTypes(apiItem *Api, linePrefix string, types map[string]TypeExpr) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) checkType(linePrefix string, types map[string]TypeExpr, expr DataType) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) readContent(filename string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *Parser) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	_ = "STUB: not implemented"
	return
}

func WithParserDebug() ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

func WithParserPrefix(prefix string) ParserOption {
	_ = "STUB: not implemented"
	return *new(ParserOption)
}

func WithParserSkipCheckTypeDeclaration() ParserOption {
	_ = "STUB: not implemented"
	return *new(ParserOption)
}
