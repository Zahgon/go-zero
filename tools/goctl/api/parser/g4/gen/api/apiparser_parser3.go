package api

import (
	"github.com/zeromicro/antlr"
)

type ITypeBlockBodyContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsTypeBlockBodyContext()
}

type TypeBlockBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeBlockBodyContext() *TypeBlockBodyContext { _ = "STUB: not implemented"; return nil }

func (*TypeBlockBodyContext) IsTypeBlockBodyContext() { _ = "STUB: not implemented"; return }

func NewTypeBlockBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockBodyContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *TypeBlockBodyContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *TypeBlockBodyContext) TypeBlockStruct() ITypeBlockStructContext {
	_ = "STUB: not implemented"
	return *new(ITypeBlockStructContext)
}

func (s *TypeBlockBodyContext) TypeBlockAlias() ITypeBlockAliasContext {
	_ = "STUB: not implemented"
	return *new(ITypeBlockAliasContext)
}

func (s *TypeBlockBodyContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *TypeBlockBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *TypeBlockBodyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) TypeBlockBody() (localctx ITypeBlockBodyContext) {
	_ = "STUB: not implemented"
	return *new(ITypeBlockBodyContext)
}

type ITypeStructContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetStructName() antlr.Token

	GetStructToken() antlr.Token

	GetLbrace() antlr.Token

	GetRbrace() antlr.Token

	SetStructName(antlr.Token)

	SetStructToken(antlr.Token)

	SetLbrace(antlr.Token)

	SetRbrace(antlr.Token)

	IsTypeStructContext()
}

type TypeStructContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	structName  antlr.Token
	structToken antlr.Token
	lbrace      antlr.Token
	rbrace      antlr.Token
}

func NewEmptyTypeStructContext() *TypeStructContext { _ = "STUB: not implemented"; return nil }

func (*TypeStructContext) IsTypeStructContext() { _ = "STUB: not implemented"; return }

func NewTypeStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeStructContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *TypeStructContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *TypeStructContext) GetStructName() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeStructContext) GetStructToken() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeStructContext) GetLbrace() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeStructContext) GetRbrace() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeStructContext) SetStructName(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeStructContext) SetStructToken(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeStructContext) SetLbrace(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeStructContext) SetRbrace(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeStructContext) AllID() []antlr.TerminalNode { _ = "STUB: not implemented"; return nil }

func (s *TypeStructContext) ID(i int) antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *TypeStructContext) AllField() []IFieldContext { _ = "STUB: not implemented"; return nil }

func (s *TypeStructContext) Field(i int) IFieldContext {
	_ = "STUB: not implemented"
	return *new(IFieldContext)
}

func (s *TypeStructContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *TypeStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *TypeStructContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) TypeStruct() (localctx ITypeStructContext) {
	_ = "STUB: not implemented"
	return *new(ITypeStructContext)
}

type ITypeAliasContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetAlias() antlr.Token

	GetAssign() antlr.Token

	SetAlias(antlr.Token)

	SetAssign(antlr.Token)

	IsTypeAliasContext()
}

type TypeAliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
	assign antlr.Token
}

func NewEmptyTypeAliasContext() *TypeAliasContext { _ = "STUB: not implemented"; return nil }

func (*TypeAliasContext) IsTypeAliasContext() { _ = "STUB: not implemented"; return }

func NewTypeAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAliasContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *TypeAliasContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *TypeAliasContext) GetAlias() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeAliasContext) GetAssign() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeAliasContext) SetAlias(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeAliasContext) SetAssign(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeAliasContext) DataType() IDataTypeContext {
	_ = "STUB: not implemented"
	return *new(IDataTypeContext)
}

func (s *TypeAliasContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *TypeAliasContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *TypeAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *TypeAliasContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) TypeAlias() (localctx ITypeAliasContext) {
	_ = "STUB: not implemented"
	return *new(ITypeAliasContext)
}

type ITypeBlockStructContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetStructName() antlr.Token

	GetStructToken() antlr.Token

	GetLbrace() antlr.Token

	GetRbrace() antlr.Token

	SetStructName(antlr.Token)

	SetStructToken(antlr.Token)

	SetLbrace(antlr.Token)

	SetRbrace(antlr.Token)

	IsTypeBlockStructContext()
}

type TypeBlockStructContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	structName  antlr.Token
	structToken antlr.Token
	lbrace      antlr.Token
	rbrace      antlr.Token
}

func NewEmptyTypeBlockStructContext() *TypeBlockStructContext {
	_ = "STUB: not implemented"
	return nil
}

func (*TypeBlockStructContext) IsTypeBlockStructContext() { _ = "STUB: not implemented"; return }

func NewTypeBlockStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockStructContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *TypeBlockStructContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *TypeBlockStructContext) GetStructName() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeBlockStructContext) GetStructToken() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeBlockStructContext) GetLbrace() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeBlockStructContext) GetRbrace() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeBlockStructContext) SetStructName(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeBlockStructContext) SetStructToken(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeBlockStructContext) SetLbrace(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeBlockStructContext) SetRbrace(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeBlockStructContext) AllID() []antlr.TerminalNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *TypeBlockStructContext) ID(i int) antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *TypeBlockStructContext) AllField() []IFieldContext { _ = "STUB: not implemented"; return nil }

func (s *TypeBlockStructContext) Field(i int) IFieldContext {
	_ = "STUB: not implemented"
	return *new(IFieldContext)
}

func (s *TypeBlockStructContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *TypeBlockStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *TypeBlockStructContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) TypeBlockStruct() (localctx ITypeBlockStructContext) {
	_ = "STUB: not implemented"
	return *new(ITypeBlockStructContext)
}
