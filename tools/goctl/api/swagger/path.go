package swagger

import (
	"github.com/go-openapi/spec"
	apiSpec "github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func spec2Paths(ctx Context, srv apiSpec.Service) *spec.Paths {
	_ = "STUB: not implemented"
	return nil
}

func mergePathItem(old, new spec.PathItem) spec.PathItem {
	_ = "STUB: not implemented"
	return *new(spec.PathItem)
}

func spec2Path(ctx Context, group apiSpec.Group, route apiSpec.Route) spec.PathItem {
	_ = "STUB: not implemented"
	return *new(spec.PathItem)
}
