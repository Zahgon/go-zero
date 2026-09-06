package api

import (
	"github.com/zeromicro/antlr"
)

func (p *ApiParserParser) DataType() (localctx IDataTypeContext) {
	_ = "STUB: not implemented"
	return *new(IDataTypeContext)
}

type IPointerTypeContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetStar() antlr.Token

	SetStar(antlr.Token)

	IsPointerTypeContext()
}

type PointerTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	star   antlr.Token
}

func NewEmptyPointerTypeContext() *PointerTypeContext { _ = "STUB: not implemented"; return nil }

func (*PointerTypeContext) IsPointerTypeContext() { _ = "STUB: not implemented"; return }

func NewPointerTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerTypeContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *PointerTypeContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *PointerTypeContext) GetStar() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *PointerTypeContext) SetStar(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *PointerTypeContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *PointerTypeContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *PointerTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *PointerTypeContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) PointerType() (localctx IPointerTypeContext) {
	_ = "STUB: not implemented"
	return *new(IPointerTypeContext)
}

type IMapTypeContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetMapToken() antlr.Token

	GetLbrack() antlr.Token

	GetKey() antlr.Token

	GetRbrack() antlr.Token

	SetMapToken(antlr.Token)

	SetLbrack(antlr.Token)

	SetKey(antlr.Token)

	SetRbrack(antlr.Token)

	GetValue() IDataTypeContext

	SetValue(IDataTypeContext)

	IsMapTypeContext()
}

type MapTypeContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	mapToken antlr.Token
	lbrack   antlr.Token
	key      antlr.Token
	rbrack   antlr.Token
	value    IDataTypeContext
}

func NewEmptyMapTypeContext() *MapTypeContext { _ = "STUB: not implemented"; return nil }

func (*MapTypeContext) IsMapTypeContext() { _ = "STUB: not implemented"; return }

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *MapTypeContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *MapTypeContext) GetMapToken() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *MapTypeContext) GetLbrack() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *MapTypeContext) GetKey() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *MapTypeContext) GetRbrack() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *MapTypeContext) SetMapToken(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *MapTypeContext) SetLbrack(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *MapTypeContext) SetKey(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *MapTypeContext) SetRbrack(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *MapTypeContext) GetValue() IDataTypeContext {
	_ = "STUB: not implemented"
	return *new(IDataTypeContext)
}

func (s *MapTypeContext) SetValue(v IDataTypeContext) { _ = "STUB: not implemented"; return }

func (s *MapTypeContext) AllID() []antlr.TerminalNode { _ = "STUB: not implemented"; return nil }

func (s *MapTypeContext) ID(i int) antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *MapTypeContext) DataType() IDataTypeContext {
	_ = "STUB: not implemented"
	return *new(IDataTypeContext)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) MapType() (localctx IMapTypeContext) {
	_ = "STUB: not implemented"
	return *new(IMapTypeContext)
}

type IArrayTypeContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetLbrack() antlr.Token

	GetRbrack() antlr.Token

	SetLbrack(antlr.Token)

	SetRbrack(antlr.Token)

	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lbrack antlr.Token
	rbrack antlr.Token
}

func NewEmptyArrayTypeContext() *ArrayTypeContext { _ = "STUB: not implemented"; return nil }

func (*ArrayTypeContext) IsArrayTypeContext() { _ = "STUB: not implemented"; return }

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ArrayTypeContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *ArrayTypeContext) GetLbrack() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *ArrayTypeContext) GetRbrack() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *ArrayTypeContext) SetLbrack(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *ArrayTypeContext) SetRbrack(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *ArrayTypeContext) DataType() IDataTypeContext {
	_ = "STUB: not implemented"
	return *new(IDataTypeContext)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) ArrayType() (localctx IArrayTypeContext) {
	_ = "STUB: not implemented"
	return *new(IArrayTypeContext)
}

type IServiceSpecContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsServiceSpecContext()
}

type ServiceSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceSpecContext() *ServiceSpecContext { _ = "STUB: not implemented"; return nil }

func (*ServiceSpecContext) IsServiceSpecContext() { _ = "STUB: not implemented"; return }

func NewServiceSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceSpecContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceSpecContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *ServiceSpecContext) ServiceApi() IServiceApiContext {
	_ = "STUB: not implemented"
	return *new(IServiceApiContext)
}

func (s *ServiceSpecContext) AtServer() IAtServerContext {
	_ = "STUB: not implemented"
	return *new(IAtServerContext)
}

func (s *ServiceSpecContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *ServiceSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ServiceSpecContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) ServiceSpec() (localctx IServiceSpecContext) {
	_ = "STUB: not implemented"
	return *new(IServiceSpecContext)
}

type IAtServerContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetLp() antlr.Token

	GetRp() antlr.Token

	SetLp(antlr.Token)

	SetRp(antlr.Token)

	IsAtServerContext()
}
