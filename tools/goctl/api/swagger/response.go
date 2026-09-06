package swagger

import (
	"github.com/go-openapi/spec"
	apiSpec "github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func jsonResponseFromType(ctx Context, atDoc apiSpec.AtDoc, tp apiSpec.Type) *spec.Responses {
	_ = "STUB: not implemented"
	return nil
}

func responsesFromStatusCode(atDoc apiSpec.AtDoc, statusCode int, response spec.Response) *spec.Responses {
	_ = "STUB: not implemented"
	return nil
}

func responseStatusCode(atDoc apiSpec.AtDoc) int { _ = "STUB: not implemented"; return 0 }

func responseDescriptions(atDoc apiSpec.AtDoc) map[int]string {
	_ = "STUB: not implemented"
	return nil
}
