package ast

import (
	"io"
	"text/tabwriter"
)

const (
	NilIndent  = ""
	WhiteSpace = " "
	Indent     = "\t"
	NewLine    = "\n"
)

const (
	_ WriteMode = 1 << iota

	ModeAuto

	ModeExpectInSameLine
)

type Option func(o *option)

type option struct {
	prefix  string
	infix   string
	mode    WriteMode
	nodes   []Node
	rawText bool
}

type tokenNodeOption func(o *tokenNodeOpt)
type tokenNodeOpt struct {
	prefix               string
	infix                string
	ignoreHeadComment    bool
	ignoreLeadingComment bool
}

type WriteMode int

type Writer struct {
	tw     *tabwriter.Writer
	writer io.Writer
}

func transfer2TokenNode(node Node, isChild bool, opt ...tokenNodeOption) *TokenNode {
	_ = "STUB: not implemented"
	return nil
}

func transferNilInfixNode(nodes []*TokenNode, opt ...tokenNodeOption) *TokenNode {
	_ = "STUB: not implemented"
	return nil
}

func transferTokenNode(node *TokenNode, opt ...tokenNodeOption) *TokenNode {
	_ = "STUB: not implemented"
	return nil
}

func ignoreHeadComment() tokenNodeOption { _ = "STUB: not implemented"; return *new(tokenNodeOption) }

func ignoreLeadingComment() tokenNodeOption {
	_ = "STUB: not implemented"
	return *new(tokenNodeOption)
}

func ignoreComment() tokenNodeOption { _ = "STUB: not implemented"; return *new(tokenNodeOption) }

func withTokenNodePrefix(prefix ...string) tokenNodeOption {
	_ = "STUB: not implemented"
	return *new(tokenNodeOption)
}

func withTokenNodeInfix(infix string) tokenNodeOption {
	_ = "STUB: not implemented"
	return *new(tokenNodeOption)
}

func expectSameLine() Option { _ = "STUB: not implemented"; return *new(Option) }

func expectIndentInfix() Option { _ = "STUB: not implemented"; return *new(Option) }

func withNode(nodes ...Node) Option { _ = "STUB: not implemented"; return *new(Option) }

func withMode(mode WriteMode) Option { _ = "STUB: not implemented"; return *new(Option) }

func withPrefix(prefix ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

func withInfix(infix string) Option { _ = "STUB: not implemented"; return *new(Option) }

func withRawText() Option { _ = "STUB: not implemented"; return *new(Option) }

func NewWriter(writer io.Writer) *Writer { _ = "STUB: not implemented"; return nil }

func NewBufferWriter() *Writer { _ = "STUB: not implemented"; return nil }

func (w *Writer) String() string { _ = "STUB: not implemented"; return "" }

func (w *Writer) Flush() { _ = "STUB: not implemented"; return }

func (w *Writer) NewLine() { _ = "STUB: not implemented"; return }

func (w *Writer) Write(opts ...Option) { _ = "STUB: not implemented"; return }

func (w *Writer) WriteText(text string) { _ = "STUB: not implemented"; return }

func (w *Writer) write(opt *option) { _ = "STUB: not implemented"; return }
