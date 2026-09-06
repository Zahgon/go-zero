package api

import (
	"github.com/zeromicro/antlr"
)

func (s *KvLitContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) KvLit() (localctx IKvLitContext) {
	_ = "STUB: not implemented"
	return *new(IKvLitContext)
}

type IServiceNameContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsServiceNameContext()
}

type ServiceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext { _ = "STUB: not implemented"; return nil }

func (*ServiceNameContext) IsServiceNameContext() { _ = "STUB: not implemented"; return }

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceNameContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *ServiceNameContext) AllID() []antlr.TerminalNode { _ = "STUB: not implemented"; return nil }

func (s *ServiceNameContext) ID(i int) antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *ServiceNameContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) ServiceName() (localctx IServiceNameContext) {
	_ = "STUB: not implemented"
	return *new(IServiceNameContext)
}

type IPathContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext { _ = "STUB: not implemented"; return nil }

func (*PathContext) IsPathContext() { _ = "STUB: not implemented"; return }

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *PathContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *PathContext) AllPathItem() []IPathItemContext { _ = "STUB: not implemented"; return nil }

func (s *PathContext) PathItem(i int) IPathItemContext {
	_ = "STUB: not implemented"
	return *new(IPathItemContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) Path() (localctx IPathContext) {
	_ = "STUB: not implemented"
	return *new(IPathContext)
}

type IPathItemContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IsPathItemContext()
}

type PathItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathItemContext() *PathItemContext { _ = "STUB: not implemented"; return nil }

func (*PathItemContext) IsPathItemContext() { _ = "STUB: not implemented"; return }

func NewPathItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathItemContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *PathItemContext) GetParser() antlr.Parser {
	_ = "STUB: not implemented"
	return *new(antlr.Parser)
}

func (s *PathItemContext) AllID() []antlr.TerminalNode { _ = "STUB: not implemented"; return nil }

func (s *PathItemContext) ID(i int) antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *PathItemContext) AllLetterOrDigit() []antlr.TerminalNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *PathItemContext) LetterOrDigit(i int) antlr.TerminalNode {
	_ = "STUB: not implemented"
	return *new(antlr.TerminalNode)
}

func (s *PathItemContext) GetRuleContext() antlr.RuleContext {
	_ = "STUB: not implemented"
	return *new(antlr.RuleContext)
}

func (s *PathItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *PathItemContext) Accept(visitor antlr.ParseTreeVisitor) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (p *ApiParserParser) PathItem() (localctx IPathItemContext) {
	_ = "STUB: not implemented"
	return *new(IPathItemContext)
}

func (p *ApiParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ApiParserParser) Field_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	_ = "STUB: not implemented"
	return false
}
