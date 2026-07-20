package swagger

import (
	"github.com/go-openapi/spec"
	apiSpec "github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func propertiesFromType(ctx Context, tp apiSpec.Type) (spec.SchemaProperties, []string) {
	_ = "STUB: not implemented"
	return *new(spec.SchemaProperties), nil
}

func containsStruct(tp apiSpec.Type) (string, bool) { _ = "STUB: not implemented"; return "", false }

func getRefName(typeName string) string { _ = "STUB: not implemented"; return "" }
