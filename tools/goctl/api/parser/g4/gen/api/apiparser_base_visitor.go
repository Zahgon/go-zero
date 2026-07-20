package api

import "github.com/zeromicro/antlr"

type BaseApiParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseApiParserVisitor) VisitApi(ctx *ApiContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitSpec(ctx *SpecContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitSyntaxLit(ctx *SyntaxLitContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitImportSpec(ctx *ImportSpecContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitImportLit(ctx *ImportLitContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitImportBlock(ctx *ImportBlockContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitImportBlockValue(ctx *ImportBlockValueContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitImportValue(ctx *ImportValueContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitInfoSpec(ctx *InfoSpecContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitTypeSpec(ctx *TypeSpecContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitTypeLit(ctx *TypeLitContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitTypeBlock(ctx *TypeBlockContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitTypeLitBody(ctx *TypeLitBodyContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitTypeBlockBody(ctx *TypeBlockBodyContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitTypeStruct(ctx *TypeStructContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitTypeAlias(ctx *TypeAliasContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitTypeBlockStruct(ctx *TypeBlockStructContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitTypeBlockAlias(ctx *TypeBlockAliasContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitField(ctx *FieldContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitNormalField(ctx *NormalFieldContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitAnonymousFiled(ctx *AnonymousFiledContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitDataType(ctx *DataTypeContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitPointerType(ctx *PointerTypeContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitMapType(ctx *MapTypeContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitArrayType(ctx *ArrayTypeContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitServiceSpec(ctx *ServiceSpecContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitAtServer(ctx *AtServerContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitServiceApi(ctx *ServiceApiContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitServiceRoute(ctx *ServiceRouteContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitAtDoc(ctx *AtDocContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitAtHandler(ctx *AtHandlerContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitRoute(ctx *RouteContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitBody(ctx *BodyContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitReplybody(ctx *ReplybodyContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitKvLit(ctx *KvLitContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitServiceName(ctx *ServiceNameContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitPath(ctx *PathContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (v *BaseApiParserVisitor) VisitPathItem(ctx *PathItemContext) any {
	_ = "STUB: not implemented"
	return *new(any)
}
