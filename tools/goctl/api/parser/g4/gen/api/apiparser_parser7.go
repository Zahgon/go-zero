package api

import (
	"github.com/zeromicro/antlr"
)

type IAtHandlerContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsAtHandlerContext()
}

type AtHandlerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtHandlerContext() *AtHandlerContext { _ = "STUB: not implemented"; return nil }

func (*AtHandlerContext) IsAtHandlerContext() { _ = "STUB: not implemented"; return }

func NewAtHandlerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtHandlerContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *AtHandlerContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *AtHandlerContext) ATHANDLER() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *AtHandlerContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *AtHandlerContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *AtHandlerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *AtHandlerContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) AtHandler() (localctx IAtHandlerContext) {
	_ = "STUB: not implemented"
	return *new(IAtHandlerContext)
}

type IRouteContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetHttpMethod() antlr.Token

	SetHttpMethod(antlr.Token)

	GetRequest() IBodyContext

	GetResponse() IReplybodyContext

	SetRequest(IBodyContext)

	SetResponse(IReplybodyContext)

	IsRouteContext()
}

type RouteContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	httpMethod antlr.Token
	request    IBodyContext
	response   IReplybodyContext
}

func NewEmptyRouteContext() *RouteContext { _ = "STUB: not implemented"; return nil }

func (*RouteContext) IsRouteContext() { _ = "STUB: not implemented"; return }

func NewRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RouteContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *RouteContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *RouteContext) GetHttpMethod() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *RouteContext) SetHttpMethod(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *RouteContext) GetRequest() IBodyContext {
	_ = "STUB: not implemented"
	return *new(IBodyContext)
}

func (s *RouteContext) GetResponse() IReplybodyContext {
	_ = "STUB: not implemented"
	return *new(IReplybodyContext)
}

func (s *RouteContext) SetRequest(v IBodyContext) { _ = "STUB: not implemented"; return }

func (s *RouteContext) SetResponse(v IReplybodyContext) { _ = "STUB: not implemented"; return }

func (s *RouteContext) Path() IPathContext { _ = "STUB: not implemented"; return *new(IPathContext) }

func (s *RouteContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *RouteContext) Body() IBodyContext { _ = "STUB: not implemented"; return *new(IBodyContext) }

func (s *RouteContext) Replybody() IReplybodyContext {
	_ = "STUB: not implemented"
	return *new(IReplybodyContext)
}

func (s *RouteContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *RouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *RouteContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) Route() (localctx IRouteContext) {
	_ = "STUB: not implemented"
	return *new(IRouteContext)
}

type IBodyContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetLp() antlr.Token

	GetRp() antlr.Token

	SetLp(antlr.Token)

	SetRp(antlr.Token)

	IsBodyContext()
}

type BodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyBodyContext() *BodyContext { _ = "STUB: not implemented"; return nil }

func (*BodyContext) IsBodyContext() { _ = "STUB: not implemented"; return }

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *BodyContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *BodyContext) GetLp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *BodyContext) GetRp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *BodyContext) SetLp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *BodyContext) SetRp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *BodyContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) Body() (localctx IBodyContext) {
	_ = "STUB: not implemented"
	return *new(IBodyContext)
}

type IReplybodyContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetReturnToken() antlr.Token

	GetLp() antlr.Token

	GetRp() antlr.Token

	SetReturnToken(antlr.Token)

	SetLp(antlr.Token)

	SetRp(antlr.Token)

	IsReplybodyContext()
}

type ReplybodyContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	returnToken antlr.Token
	lp          antlr.Token
	rp          antlr.Token
}

func NewEmptyReplybodyContext() *ReplybodyContext { _ = "STUB: not implemented"; return nil }

func (*ReplybodyContext) IsReplybodyContext() { _ = "STUB: not implemented"; return }

func NewReplybodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplybodyContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReplybodyContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *ReplybodyContext) GetReturnToken() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *ReplybodyContext) GetLp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *ReplybodyContext) GetRp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *ReplybodyContext) SetReturnToken(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *ReplybodyContext) SetLp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *ReplybodyContext) SetRp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *ReplybodyContext) DataType() IDataTypeContext {
	_ = "STUB: not implemented"
	return *new(IDataTypeContext)
}

func (s *ReplybodyContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *ReplybodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ReplybodyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) Replybody() (localctx IReplybodyContext) {
	_ = "STUB: not implemented"
	return *new(IReplybodyContext)
}

type IKvLitContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetKey() antlr.Token

	GetValue() antlr.Token

	SetKey(antlr.Token)

	SetValue(antlr.Token)

	IsKvLitContext()
}

type KvLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  antlr.Token
}

func NewEmptyKvLitContext() *KvLitContext { _ = "STUB: not implemented"; return nil }

func (*KvLitContext) IsKvLitContext() { _ = "STUB: not implemented"; return }

func NewKvLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KvLitContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *KvLitContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *KvLitContext) GetKey() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *KvLitContext) GetValue() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *KvLitContext) SetKey(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *KvLitContext) SetValue(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *KvLitContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *KvLitContext) LINE_VALUE() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *KvLitContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *KvLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}
