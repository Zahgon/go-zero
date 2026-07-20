package api

import (
	"github.com/zeromicro/antlr"
)

type ITypeBlockAliasContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetAlias() antlr.Token

	GetAssign() antlr.Token

	SetAlias(antlr.Token)

	SetAssign(antlr.Token)

	IsTypeBlockAliasContext()
}

type TypeBlockAliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
	assign antlr.Token
}

func NewEmptyTypeBlockAliasContext() *TypeBlockAliasContext { _ = "STUB: not implemented"; return nil }

func (*TypeBlockAliasContext) IsTypeBlockAliasContext() { _ = "STUB: not implemented"; return }

func NewTypeBlockAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockAliasContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *TypeBlockAliasContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *TypeBlockAliasContext) GetAlias() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeBlockAliasContext) GetAssign() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *TypeBlockAliasContext) SetAlias(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeBlockAliasContext) SetAssign(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *TypeBlockAliasContext) DataType() IDataTypeContext {
	_ = "STUB: not implemented"
	return *new(IDataTypeContext)
}

func (s *TypeBlockAliasContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *TypeBlockAliasContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *TypeBlockAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *TypeBlockAliasContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) TypeBlockAlias() (localctx ITypeBlockAliasContext) {
	_ = "STUB: not implemented"
	return *new(ITypeBlockAliasContext)
}

type IFieldContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext { _ = "STUB: not implemented"; return nil }

func (*FieldContext) IsFieldContext() { _ = "STUB: not implemented"; return }

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *FieldContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *FieldContext) NormalField() INormalFieldContext {
	_ = "STUB: not implemented"
	return *new(INormalFieldContext)
}

func (s *FieldContext) AnonymousFiled() IAnonymousFiledContext {
	_ = "STUB: not implemented"
	return *new(IAnonymousFiledContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) Field() (localctx IFieldContext) {
	_ = "STUB: not implemented"
	return *new(IFieldContext)
}

type INormalFieldContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetFieldName() antlr.Token

	GetTag() antlr.Token

	SetFieldName(antlr.Token)

	SetTag(antlr.Token)

	IsNormalFieldContext()
}

type NormalFieldContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	fieldName antlr.Token
	tag       antlr.Token
}

func NewEmptyNormalFieldContext() *NormalFieldContext { _ = "STUB: not implemented"; return nil }

func (*NormalFieldContext) IsNormalFieldContext() { _ = "STUB: not implemented"; return }

func NewNormalFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NormalFieldContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *NormalFieldContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *NormalFieldContext) GetFieldName() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *NormalFieldContext) GetTag() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *NormalFieldContext) SetFieldName(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *NormalFieldContext) SetTag(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *NormalFieldContext) DataType() IDataTypeContext {
	_ = "STUB: not implemented"
	return *new(IDataTypeContext)
}

func (s *NormalFieldContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *NormalFieldContext) RAW_STRING() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *NormalFieldContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *NormalFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *NormalFieldContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) NormalField() (localctx INormalFieldContext) {
	_ = "STUB: not implemented"
	return *new(INormalFieldContext)
}

type IAnonymousFiledContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetStar() antlr.Token

	SetStar(antlr.Token)

	IsAnonymousFiledContext()
}

type AnonymousFiledContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	star   antlr.Token
}

func NewEmptyAnonymousFiledContext() *AnonymousFiledContext { _ = "STUB: not implemented"; return nil }

func (*AnonymousFiledContext) IsAnonymousFiledContext() { _ = "STUB: not implemented"; return }

func NewAnonymousFiledContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnonymousFiledContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *AnonymousFiledContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *AnonymousFiledContext) GetStar() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *AnonymousFiledContext) SetStar(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *AnonymousFiledContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *AnonymousFiledContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *AnonymousFiledContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *AnonymousFiledContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) AnonymousFiled() (localctx IAnonymousFiledContext) {
	_ = "STUB: not implemented"
	return *new(IAnonymousFiledContext)
}

type IDataTypeContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetInter() antlr.Token

	GetTime() antlr.Token

	SetInter(antlr.Token)

	SetTime(antlr.Token)

	IsDataTypeContext()
}

type DataTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	inter  antlr.Token
	time   antlr.Token
}

func NewEmptyDataTypeContext() *DataTypeContext { _ = "STUB: not implemented"; return nil }

func (*DataTypeContext) IsDataTypeContext() { _ = "STUB: not implemented"; return }

func NewDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *DataTypeContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *DataTypeContext) GetInter() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *DataTypeContext) GetTime() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *DataTypeContext) SetInter(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *DataTypeContext) SetTime(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *DataTypeContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *DataTypeContext) MapType() IMapTypeContext {
	_ = "STUB: not implemented"
	return *new(IMapTypeContext)
}

func (s *DataTypeContext) ArrayType() IArrayTypeContext {
	_ = "STUB: not implemented"
	return *new(IArrayTypeContext)
}

func (s *DataTypeContext) INTERFACE() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *DataTypeContext) PointerType() IPointerTypeContext {
	_ = "STUB: not implemented"
	return *new(IPointerTypeContext)
}

func (s *DataTypeContext) TypeStruct() ITypeStructContext {
	_ = "STUB: not implemented"
	return *new(ITypeStructContext)
}

func (s *DataTypeContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *DataTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *DataTypeContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}
