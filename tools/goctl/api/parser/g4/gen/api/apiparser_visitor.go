package api

import "github.com/zeromicro/antlr"

type ApiParserVisitor interface {
	antlr.ParseTreeVisitor

	VisitApi(ctx *ApiContext) any

	VisitSpec(ctx *SpecContext) any

	VisitSyntaxLit(ctx *SyntaxLitContext) any

	VisitImportSpec(ctx *ImportSpecContext) any

	VisitImportLit(ctx *ImportLitContext) any

	VisitImportBlock(ctx *ImportBlockContext) any

	VisitImportBlockValue(ctx *ImportBlockValueContext) any

	VisitImportValue(ctx *ImportValueContext) any

	VisitInfoSpec(ctx *InfoSpecContext) any

	VisitTypeSpec(ctx *TypeSpecContext) any

	VisitTypeLit(ctx *TypeLitContext) any

	VisitTypeBlock(ctx *TypeBlockContext) any

	VisitTypeLitBody(ctx *TypeLitBodyContext) any

	VisitTypeBlockBody(ctx *TypeBlockBodyContext) any

	VisitTypeStruct(ctx *TypeStructContext) any

	VisitTypeAlias(ctx *TypeAliasContext) any

	VisitTypeBlockStruct(ctx *TypeBlockStructContext) any

	VisitTypeBlockAlias(ctx *TypeBlockAliasContext) any

	VisitField(ctx *FieldContext) any

	VisitNormalField(ctx *NormalFieldContext) any

	VisitAnonymousFiled(ctx *AnonymousFiledContext) any

	VisitDataType(ctx *DataTypeContext) any

	VisitPointerType(ctx *PointerTypeContext) any

	VisitMapType(ctx *MapTypeContext) any

	VisitArrayType(ctx *ArrayTypeContext) any

	VisitServiceSpec(ctx *ServiceSpecContext) any

	VisitAtServer(ctx *AtServerContext) any

	VisitServiceApi(ctx *ServiceApiContext) any

	VisitServiceRoute(ctx *ServiceRouteContext) any

	VisitAtDoc(ctx *AtDocContext) any

	VisitAtHandler(ctx *AtHandlerContext) any

	VisitRoute(ctx *RouteContext) any

	VisitBody(ctx *BodyContext) any

	VisitReplybody(ctx *ReplybodyContext) any

	VisitKvLit(ctx *KvLitContext) any

	VisitServiceName(ctx *ServiceNameContext) any

	VisitPath(ctx *PathContext) any

	VisitPathItem(ctx *PathItemContext) any
}
