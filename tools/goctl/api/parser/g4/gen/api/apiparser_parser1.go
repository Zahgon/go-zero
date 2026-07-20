package api

import (
	"github.com/zeromicro/antlr"
)

func (s *SyntaxLitContext) STRING() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *SyntaxLitContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *SyntaxLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SyntaxLitContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) SyntaxLit() (localctx ISyntaxLitContext) {
	_ = "STUB: not implemented"
	return *new(ISyntaxLitContext)
}

type IImportSpecContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsImportSpecContext()
}

type ImportSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportSpecContext() *ImportSpecContext { _ = "STUB: not implemented"; return nil }

func (*ImportSpecContext) IsImportSpecContext() { _ = "STUB: not implemented"; return }

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ImportSpecContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *ImportSpecContext) ImportLit() IImportLitContext {
	_ = "STUB: not implemented"
	return *new(IImportLitContext)
}

func (s *ImportSpecContext) ImportBlock() IImportBlockContext {
	_ = "STUB: not implemented"
	return *new(IImportBlockContext)
}

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) ImportSpec() (localctx IImportSpecContext) {
	_ = "STUB: not implemented"
	return *new(IImportSpecContext)
}

type IImportLitContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetImportToken() antlr.Token

	SetImportToken(antlr.Token)

	IsImportLitContext()
}

type ImportLitContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	importToken antlr.Token
}

func NewEmptyImportLitContext() *ImportLitContext { _ = "STUB: not implemented"; return nil }

func (*ImportLitContext) IsImportLitContext() { _ = "STUB: not implemented"; return }

func NewImportLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportLitContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ImportLitContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *ImportLitContext) GetImportToken() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *ImportLitContext) SetImportToken(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *ImportLitContext) ImportValue() IImportValueContext {
	_ = "STUB: not implemented"
	return *new(IImportValueContext)
}

func (s *ImportLitContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *ImportLitContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *ImportLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ImportLitContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) ImportLit() (localctx IImportLitContext) {
	_ = "STUB: not implemented"
	return *new(IImportLitContext)
}

type IImportBlockContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetImportToken() antlr.Token

	SetImportToken(antlr.Token)

	IsImportBlockContext()
}

type ImportBlockContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	importToken antlr.Token
}

func NewEmptyImportBlockContext() *ImportBlockContext { _ = "STUB: not implemented"; return nil }

func (*ImportBlockContext) IsImportBlockContext() { _ = "STUB: not implemented"; return }

func NewImportBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportBlockContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ImportBlockContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *ImportBlockContext) GetImportToken() antlr.Token {
	_ = "STUB: not implemented"
	return *new(antlr.Token)
}

func (s *ImportBlockContext) SetImportToken(v antlr.Token) { _ = "STUB: not implemented"; return }

func (s *ImportBlockContext) ID() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *ImportBlockContext) AllImportBlockValue() []IImportBlockValueContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ImportBlockContext) ImportBlockValue(i int) IImportBlockValueContext {
	_ = "STUB: not implemented"
	return *new(IImportBlockValueContext)
}

func (s *ImportBlockContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *ImportBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ImportBlockContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) ImportBlock() (localctx IImportBlockContext) {
	_ = "STUB: not implemented"
	return *new(IImportBlockContext)
}

type IImportBlockValueContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsImportBlockValueContext()
}

type ImportBlockValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportBlockValueContext() *ImportBlockValueContext {
	_ = "STUB: not implemented"
	return nil
}

func (*ImportBlockValueContext) IsImportBlockValueContext() { _ = "STUB: not implemented"; return }

func NewImportBlockValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportBlockValueContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ImportBlockValueContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *ImportBlockValueContext) ImportValue() IImportValueContext {
	_ = "STUB: not implemented"
	return *new(IImportValueContext)
}

func (s *ImportBlockValueContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *ImportBlockValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ImportBlockValueContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) ImportBlockValue() (localctx IImportBlockValueContext) {
	_ = "STUB: not implemented"
	return *new(IImportBlockValueContext)
}

type IImportValueContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsImportValueContext()
}

type ImportValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportValueContext() *ImportValueContext { _ = "STUB: not implemented"; return nil }

func (*ImportValueContext) IsImportValueContext() { _ = "STUB: not implemented"; return }

func NewImportValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportValueContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ImportValueContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *ImportValueContext) STRING() antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *ImportValueContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *ImportValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ImportValueContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) ImportValue() (localctx IImportValueContext) {
	_ = "STUB: not implemented"
	return *new(IImportValueContext)
}

type IInfoSpecContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	GetInfoToken() antlr.Token

	GetLp() antlr.Token

	GetRp() antlr.Token

	SetInfoToken(antlr.Token)

	SetLp(antlr.Token)

	SetRp(antlr.Token)

	IsInfoSpecContext()
}
