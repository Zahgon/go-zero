package gogen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

var (
	//go:embed logic.tpl
	logicTemplate string

	//go:embed sse_logic.tpl
	sseLogicTemplate string
)

func genLogic(dir, rootPkg, projectPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func genLogicByRoute(dir, rootPkg, projectPkg string, cfg *config.Config, group spec.Group, route spec.Route) error {
	_ = "STUB: not implemented"
	return nil
}

func getLogicFolderPath(group spec.Group, route spec.Route) string {
	_ = "STUB: not implemented"
	return ""
}

func genLogicImports(route spec.Route, parentPkg string) string {
	_ = "STUB: not implemented"
	return ""
}

func onlyPrimitiveTypes(val string) bool { _ = "STUB: not implemented"; return false }

func shallImportTypesPackage(route spec.Route) bool { _ = "STUB: not implemented"; return false }
