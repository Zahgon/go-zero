package api

import (
	"github.com/zeromicro/antlr"
)

type AtServerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyAtServerContext() *AtServerContext { _ = "STUB: not implemented"; return nil }

func (*AtServerContext) IsAtServerContext() { _ = "STUB: not implemented"; return }

func NewAtServerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtServerContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *AtServerContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *AtServerContext) GetLp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *AtServerContext) GetRp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *AtServerContext) SetLp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *AtServerContext) SetRp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *AtServerContext) ATSERVER() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *AtServerContext) AllKvLit() []IKvLitContext { _ = "STUB: not implemented"; return nil }

func (s *AtServerContext) KvLit(i int) IKvLitContext {
	_ = "STUB: not implemented"
	return *new(IKvLitContext)
}

func (s *AtServerContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *AtServerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *AtServerContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) AtServer() (localctx IAtServerContext) {
	_ = "STUB: not implemented"
	return *new(IAtServerContext)
}

type IServiceApiContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetServiceToken() antlr.Token

	GetLbrace() antlr.Token

	GetRbrace() antlr.Token

	SetServiceToken(antlr.Token)

	SetLbrace(antlr.Token)

	SetRbrace(antlr.Token)

	IsServiceApiContext()
}

type ServiceApiContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	serviceToken antlr.Token
	lbrace       antlr.Token
	rbrace       antlr.Token
}

func NewEmptyServiceApiContext() *ServiceApiContext { _ = "STUB: not implemented"; return nil }

func (*ServiceApiContext) IsServiceApiContext() { _ = "STUB: not implemented"; return }

func NewServiceApiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceApiContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceApiContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *ServiceApiContext) GetServiceToken() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *ServiceApiContext) GetLbrace() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *ServiceApiContext) GetRbrace() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *ServiceApiContext) SetServiceToken(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *ServiceApiContext) SetLbrace(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *ServiceApiContext) SetRbrace(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *ServiceApiContext) ServiceName() IServiceNameContext {
	_ = "STUB: not implemented"
	return *new(IServiceNameContext)
}

func (s *ServiceApiContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *ServiceApiContext) AllServiceRoute() []IServiceRouteContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceApiContext) ServiceRoute(i int) IServiceRouteContext {
	_ = "STUB: not implemented"
	return *new(IServiceRouteContext)
}

func (s *ServiceApiContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *ServiceApiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ServiceApiContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) ServiceApi() (localctx IServiceApiContext) {
	_ = "STUB: not implemented"
	return *new(IServiceApiContext)
}

type IServiceRouteContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsServiceRouteContext()
}

type ServiceRouteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceRouteContext() *ServiceRouteContext { _ = "STUB: not implemented"; return nil }

func (*ServiceRouteContext) IsServiceRouteContext() { _ = "STUB: not implemented"; return }

func NewServiceRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceRouteContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceRouteContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *ServiceRouteContext) Route() IRouteContext {
	_ = "STUB: not implemented"
	return *new(IRouteContext)
}

func (s *ServiceRouteContext) AtServer() IAtServerContext {
	_ = "STUB: not implemented"
	return *new(IAtServerContext)
}

func (s *ServiceRouteContext) AtHandler() IAtHandlerContext {
	_ = "STUB: not implemented"
	return *new(IAtHandlerContext)
}

func (s *ServiceRouteContext) AtDoc() IAtDocContext {
	_ = "STUB: not implemented"
	return *new(IAtDocContext)
}

func (s *ServiceRouteContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *ServiceRouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ServiceRouteContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) ServiceRoute() (localctx IServiceRouteContext) {
	_ = "STUB: not implemented"
	return *new(IServiceRouteContext)
}

type IAtDocContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetLp() antlr.Token

	GetRp() antlr.Token

	SetLp(antlr.Token)

	SetRp(antlr.Token)

	IsAtDocContext()
}

type AtDocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyAtDocContext() *AtDocContext { _ = "STUB: not implemented"; return nil }

func (*AtDocContext) IsAtDocContext() { _ = "STUB: not implemented"; return }

func NewAtDocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtDocContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *AtDocContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *AtDocContext) GetLp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *AtDocContext) GetRp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *AtDocContext) SetLp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *AtDocContext) SetRp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *AtDocContext) ATDOC() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *AtDocContext) STRING() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *AtDocContext) AllKvLit() []IKvLitContext { _ = "STUB: not implemented"; return nil }

func (s *AtDocContext) KvLit(i int) IKvLitContext {
	_ = "STUB: not implemented"
	return *new(IKvLitContext)
}

func (s *AtDocContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *AtDocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *AtDocContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) AtDoc() (localctx IAtDocContext) {
	_ = "STUB: not implemented"
	return *new(IAtDocContext)
}
