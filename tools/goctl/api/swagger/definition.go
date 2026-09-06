package swagger

import (
	"github.com/go-openapi/spec"
	apiSpec "github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func definitionsFromTypes(ctx Context, types []apiSpec.Type) spec.Definitions {
	_ = "STUB: not implemented"
	return *new(spec.Definitions)
}

func schemaFromType(ctx Context, tp apiSpec.Type) spec.Schema {
	_ = "STUB: not implemented"
	return *new(spec.Schema)
}
