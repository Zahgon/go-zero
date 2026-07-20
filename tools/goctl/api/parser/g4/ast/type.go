package ast

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"
)

type (
	TypeExpr interface {
		Doc() []Expr
		Format() error
		Equal(v any) bool
		NameExpr() Expr
	}

	TypeAlias struct {
		Name        Expr
		Assign      Expr
		DataType    DataType
		DocExpr     []Expr
		CommentExpr Expr
	}

	TypeStruct struct {
		Name    Expr
		Struct  Expr
		LBrace  Expr
		RBrace  Expr
		DocExpr []Expr
		Fields  []*TypeField
	}

	TypeField struct {
		IsAnonymous bool

		Name        Expr
		DataType    DataType
		Tag         Expr
		DocExpr     []Expr
		CommentExpr Expr
	}

	DataType interface {
		Expr() Expr
		Equal(dt DataType) bool
		Format() error
		IsNotNil() bool
	}

	Literal struct {
		Literal Expr
	}

	Interface struct {
		Literal Expr
	}

	Map struct {
		MapExpr Expr
		Map     Expr
		LBrack  Expr
		RBrack  Expr
		Key     Expr
		Value   DataType
	}

	Array struct {
		ArrayExpr Expr
		LBrack    Expr
		RBrack    Expr
		Literal   DataType
	}

	Time struct {
		Literal Expr
	}

	Pointer struct {
		PointerExpr Expr
		Star        Expr
		Name        Expr
	}
)

func (v *ApiVisitor) VisitTypeSpec(ctx *api.TypeSpecContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitTypeLit(ctx *api.TypeLitContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitTypeBlock(ctx *api.TypeBlockContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitTypeLitBody(ctx *api.TypeLitBodyContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitTypeBlockBody(ctx *api.TypeBlockBodyContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitTypeStruct(ctx *api.TypeStructContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitTypeBlockStruct(ctx *api.TypeBlockStructContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitTypeBlockAlias(ctx *api.TypeBlockAliasContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitTypeAlias(ctx *api.TypeAliasContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitField(ctx *api.FieldContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitNormalField(ctx *api.NormalFieldContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitAnonymousFiled(ctx *api.AnonymousFiledContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitDataType(ctx *api.DataTypeContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitPointerType(ctx *api.PointerTypeContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitMapType(ctx *api.MapTypeContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *ApiVisitor) VisitArrayType(ctx *api.ArrayTypeContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (a *TypeAlias) NameExpr() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (a *TypeAlias) Doc() []Expr { _ = "STUB: not implemented"; return nil }

func (a *TypeAlias) Comment() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (a *TypeAlias) Format() error { _ = "STUB: not implemented"; return nil }

func (a *TypeAlias) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (l *Literal) Expr() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (l *Literal) Format() error { _ = "STUB: not implemented"; return nil }

func (l *Literal) Equal(dt DataType) bool { _ = "STUB: not implemented"; return false }

func (l *Literal) IsNotNil() bool { _ = "STUB: not implemented"; return false }

func (i *Interface) Expr() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (i *Interface) Format() error { _ = "STUB: not implemented"; return nil }

func (i *Interface) Equal(dt DataType) bool { _ = "STUB: not implemented"; return false }

func (i *Interface) IsNotNil() bool { _ = "STUB: not implemented"; return false }

func (m *Map) Expr() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (m *Map) Format() error { _ = "STUB: not implemented"; return nil }

func (m *Map) Equal(dt DataType) bool { _ = "STUB: not implemented"; return false }

func (m *Map) IsNotNil() bool { _ = "STUB: not implemented"; return false }

func (a *Array) Expr() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (a *Array) Format() error { _ = "STUB: not implemented"; return nil }

func (a *Array) Equal(dt DataType) bool { _ = "STUB: not implemented"; return false }

func (a *Array) IsNotNil() bool { _ = "STUB: not implemented"; return false }

func (t *Time) Expr() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (t *Time) Format() error { _ = "STUB: not implemented"; return nil }

func (t *Time) Equal(dt DataType) bool { _ = "STUB: not implemented"; return false }

func (t *Time) IsNotNil() bool { _ = "STUB: not implemented"; return false }

func (p *Pointer) Expr() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (p *Pointer) Format() error { _ = "STUB: not implemented"; return nil }

func (p *Pointer) Equal(dt DataType) bool { _ = "STUB: not implemented"; return false }

func (p *Pointer) IsNotNil() bool { _ = "STUB: not implemented"; return false }

func (s *TypeStruct) NameExpr() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (s *TypeStruct) Equal(dt any) bool { _ = "STUB: not implemented"; return false }

func (s *TypeStruct) Doc() []Expr { _ = "STUB: not implemented"; return nil }

func (s *TypeStruct) Format() error { _ = "STUB: not implemented"; return nil }

func (t *TypeField) Equal(v any) bool { _ = "STUB: not implemented"; return false }

func (t *TypeField) Doc() []Expr { _ = "STUB: not implemented"; return nil }

func (t *TypeField) Comment() Expr { _ = "STUB: not implemented"; return *new(Expr) }

func (t *TypeField) Format() error { _ = "STUB: not implemented"; return nil }
