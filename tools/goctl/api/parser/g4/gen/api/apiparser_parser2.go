package api

import (
	"github.com/zeromicro/antlr"
)

type InfoSpecContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	infoToken antlr.Token
	lp        antlr.Token
	rp        antlr.Token
}

func NewEmptyInfoSpecContext() *InfoSpecContext { _ = "STUB: not implemented"; return nil }

func (*InfoSpecContext) IsInfoSpecContext() { _ = "STUB: not implemented"; return }

func NewInfoSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfoSpecContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *InfoSpecContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *InfoSpecContext) GetInfoToken() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *InfoSpecContext) GetLp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *InfoSpecContext) GetRp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *InfoSpecContext) SetInfoToken(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *InfoSpecContext) SetLp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *InfoSpecContext) SetRp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *InfoSpecContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *InfoSpecContext) AllKvLit() []IKvLitContext { _ = "STUB: not implemented"; return nil }

func (s *InfoSpecContext) KvLit(i int) IKvLitContext {
	_ = "STUB: not implemented"
	return *new(IKvLitContext)
}

func (s *InfoSpecContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *InfoSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *InfoSpecContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) InfoSpec() (localctx IInfoSpecContext) {
	_ = "STUB: not implemented"
	return *new(IInfoSpecContext)
}

type ITypeSpecContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext { _ = "STUB: not implemented"; return nil }

func (*TypeSpecContext) IsTypeSpecContext() { _ = "STUB: not implemented"; return }

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *TypeSpecContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *TypeSpecContext) TypeLit() ITypeLitContext {
	_ = "STUB: not implemented"
	return *new(ITypeLitContext)
}

func (s *TypeSpecContext) TypeBlock() ITypeBlockContext {
	_ = "STUB: not implemented"
	return *new(ITypeBlockContext)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) TypeSpec() (localctx ITypeSpecContext) {
	_ = "STUB: not implemented"
	return *new(ITypeSpecContext)
}

type ITypeLitContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetTypeToken() antlr.Token

	SetTypeToken(antlr.Token)

	IsTypeLitContext()
}

type TypeLitContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	typeToken antlr.Token
}

func NewEmptyTypeLitContext() *TypeLitContext { _ = "STUB: not implemented"; return nil }

func (*TypeLitContext) IsTypeLitContext() { _ = "STUB: not implemented"; return }

func NewTypeLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLitContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *TypeLitContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *TypeLitContext) GetTypeToken() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeLitContext) SetTypeToken(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeLitContext) TypeLitBody() ITypeLitBodyContext {
	_ = "STUB: not implemented"
	return *new(ITypeLitBodyContext)
}

func (s *TypeLitContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *TypeLitContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *TypeLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *TypeLitContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) TypeLit() (localctx ITypeLitContext) {
	_ = "STUB: not implemented"
	return *new(ITypeLitContext)
}

type ITypeBlockContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetTypeToken() antlr.Token

	GetLp() antlr.Token

	GetRp() antlr.Token

	SetTypeToken(antlr.Token)

	SetLp(antlr.Token)

	SetRp(antlr.Token)

	IsTypeBlockContext()
}

type TypeBlockContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	typeToken antlr.Token
	lp        antlr.Token
	rp        antlr.Token
}

func NewEmptyTypeBlockContext() *TypeBlockContext { _ = "STUB: not implemented"; return nil }

func (*TypeBlockContext) IsTypeBlockContext() { _ = "STUB: not implemented"; return }

func NewTypeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *TypeBlockContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *TypeBlockContext) GetTypeToken() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeBlockContext) GetLp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *TypeBlockContext) GetRp() antlr.Token { _ = "STUB: not implemented"; return *new(antlr.Token) }

func (s *TypeBlockContext) SetTypeToken(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeBlockContext) SetLp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeBlockContext) SetRp(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeBlockContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *TypeBlockContext) AllTypeBlockBody() []ITypeBlockBodyContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *TypeBlockContext) TypeBlockBody(i int) ITypeBlockBodyContext {
	_ = "STUB: not implemented"
	return *new(ITypeBlockBodyContext)
}

func (s *TypeBlockContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *TypeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *TypeBlockContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) TypeBlock() (localctx ITypeBlockContext) {
	_ = "STUB: not implemented"
	return *new(ITypeBlockContext)
}

type ITypeLitBodyContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsTypeLitBodyContext()
}

type TypeLitBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeLitBodyContext() *TypeLitBodyContext { _ = "STUB: not implemented"; return nil }

func (*TypeLitBodyContext) IsTypeLitBodyContext() { _ = "STUB: not implemented"; return }

func NewTypeLitBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLitBodyContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *TypeLitBodyContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *TypeLitBodyContext) TypeStruct() ITypeStructContext {
	_ = "STUB: not implemented"
	return *new(ITypeStructContext)
}

func (s *TypeLitBodyContext) TypeAlias() ITypeAliasContext {
	_ = "STUB: not implemented"
	return *new(ITypeAliasContext)
}

func (s *TypeLitBodyContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *TypeLitBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *TypeLitBodyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) TypeLitBody() (localctx ITypeLitBodyContext) {
	_ = "STUB: not implemented"
	return *new(ITypeLitBodyContext)
}
