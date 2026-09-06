package swagger

import (
	"github.com/go-openapi/spec"
	apiSpec "github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func isRequestBodyJson(ctx Context, method string, tp apiSpec.Type) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func parametersFromType(ctx Context, method string, tp apiSpec.Type) []spec.Parameter {
	_ = "STUB: not implemented"
	return nil
}
