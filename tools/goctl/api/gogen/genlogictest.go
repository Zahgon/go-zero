package gogen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

//go:embed logic_test.tpl
var logicTestTemplate string

func genLogicTest(dir, rootPkg, projectPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func genLogicTestByRoute(dir, rootPkg, projectPkg string, cfg *config.Config, group spec.Group, route spec.Route) error {
	_ = "STUB: not implemented"
	return nil
}

func genLogicTestImports(route spec.Route, parentPkg string) string {
	_ = "STUB: not implemented"
	return ""
}
