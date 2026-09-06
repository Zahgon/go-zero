package ast

import (
	"github.com/zeromicro/antlr"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
)

type (
	TokenStream interface {
		GetStart() antlr.Token
		GetStop() antlr.Token
		GetParser() antlr.Parser
	}

	ApiVisitor struct {
		*api.BaseApiParserVisitor
		debug    bool
		log      console.Console
		prefix   string
		infoFlag bool
	}

	VisitorOption func(v *ApiVisitor)

	Spec interface {
		Doc() []Expr
		Comment() Expr
		Format() error
		Equal(v any) bool
	}

	Expr interface {
		Prefix() string
		Line() int
		Column() int
		Text() string
		SetText(text string)
		Start() int
		Stop() int
		Equal(expr Expr) bool
		IsNotNil() bool
	}
)

func NewApiVisitor(options ...VisitorOption) *ApiVisitor { _ = "STUB: not implemented"; return nil }

func (v *ApiVisitor) panic(expr Expr, msg string) { _ = "STUB: not implemented"; return }

func WithVisitorPrefix(prefix string) VisitorOption {
	_ = "STUB: not implemented"
	return *new(VisitorOption)
}

func WithVisitorDebug() VisitorOption { _ = "STUB: not implemented"; return *new(VisitorOption) }

type defaultExpr struct {
	prefix, v    string
	line, column int
	start, stop  int
}

func NewTextExpr(v string) *defaultExpr { _ = "STUB: not implemented"; return nil }

func (v *ApiVisitor) newExprWithTerminalNode(node antlr.TerminalNode) *defaultExpr {
	_ = "STUB: not implemented"
	return nil
}

func (v *ApiVisitor) newExprWithToken(token antlr.Token) *defaultExpr {
	_ = "STUB: not implemented"
	return nil
}

func (v *ApiVisitor) newExprWithText(text string, line, column, start, stop int) *defaultExpr {
	_ = "STUB: not implemented"
	return nil
}

func (e *defaultExpr) Prefix() string { _ = "STUB: not implemented"; return "" }

func (e *defaultExpr) Line() int { _ = "STUB: not implemented"; return 0 }

func (e *defaultExpr) Column() int { _ = "STUB: not implemented"; return 0 }

func (e *defaultExpr) Text() string { _ = "STUB: not implemented"; return "" }

func (e *defaultExpr) SetText(text string) { _ = "STUB: not implemented"; return }

func (e *defaultExpr) Start() int { _ = "STUB: not implemented"; return 0 }

func (e *defaultExpr) Stop() int { _ = "STUB: not implemented"; return 0 }

func (e *defaultExpr) Equal(expr Expr) bool { _ = "STUB: not implemented"; return false }

func (e *defaultExpr) IsNotNil() bool { _ = "STUB: not implemented"; return false }

func EqualDoc(spec1, spec2 Spec) bool { _ = "STUB: not implemented"; return false }

func (v *ApiVisitor) getDoc(t TokenStream) []Expr { _ = "STUB: not implemented"; return nil }

func (v *ApiVisitor) getComment(t TokenStream) Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (v *ApiVisitor) getHiddenTokensToLeft(t TokenStream, channel int, containsCommentOfDefaultChannel bool) []Expr {
	_ = "STUB: not implemented"
	return nil
}

func (v *ApiVisitor) getHiddenTokensToRight(t TokenStream, channel int) []Expr {
	_ = "STUB: not implemented"
	return nil
}
