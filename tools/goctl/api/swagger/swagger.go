package swagger

import (
	"github.com/go-openapi/spec"
	apiSpec "github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func spec2Swagger(api *apiSpec.ApiSpec) (*spec.Swagger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatComment(comment string) string { _ = "STUB: not implemented"; return "" }

func sampleItemsFromGoType(ctx Context, tp apiSpec.Type) *spec.Items {
	_ = "STUB: not implemented"
	return nil
}

func itemsFromGoType(ctx Context, tp apiSpec.Type) *spec.SchemaOrArray {
	_ = "STUB: not implemented"
	return nil
}

func mapFromGoType(ctx Context, tp apiSpec.Type) *spec.SchemaOrBool {
	_ = "STUB: not implemented"
	return nil
}

func itemFromGoType(ctx Context, tp apiSpec.Type) *spec.SchemaOrArray {
	_ = "STUB: not implemented"
	return nil
}

func typeFromGoType(ctx Context, tp apiSpec.Type) []string { _ = "STUB: not implemented"; return nil }

func sampleTypeFromGoType(ctx Context, tp apiSpec.Type) string {
	_ = "STUB: not implemented"
	return ""
}

func typeContainsTag(ctx Context, structType apiSpec.DefineStruct, tag string) bool {
	_ = "STUB: not implemented"
	return false
}

func expandMembers(ctx Context, tp apiSpec.Type) []apiSpec.Member {
	_ = "STUB: not implemented"
	return nil
}

func rangeMemberAndDo(ctx Context, structType apiSpec.Type, do func(tag *apiSpec.Tags, required bool, member apiSpec.Member)) {
	_ = "STUB: not implemented"
	return
}

func isRequired(ctx Context, tags *apiSpec.Tags) bool { _ = "STUB: not implemented"; return false }

func isOptional(_ Context, options []string) bool { _ = "STUB: not implemented"; return false }

func pathVariable2SwaggerVariable(_ Context, path string) string {
	_ = "STUB: not implemented"
	return ""
}

func wrapCodeMsgProps(ctx Context, properties spec.SchemaProps, atDoc apiSpec.AtDoc) spec.SchemaProps {
	_ = "STUB: not implemented"
	return *new(spec.SchemaProps)
}

func specExtensions(api apiSpec.Info) (spec.Extensions, *spec.Info) {
	_ = "STUB: not implemented"
	return *new(spec.Extensions), nil
}
