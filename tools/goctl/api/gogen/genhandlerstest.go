package gogen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

//go:embed handler_test.tpl
var handlerTestTemplate string

func genHandlerTest(dir, rootPkg, projectPkg string, cfg *config.Config, group spec.Group, route spec.Route) error {
	_ = "STUB: not implemented"
	return nil
}

func genHandlersTest(dir, rootPkg, projectPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func genHandlerTestImports(group spec.Group, route spec.Route, parentPkg string) string {
	_ = "STUB: not implemented"
	return ""
}
