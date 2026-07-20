package gogen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

const defaultLogicPackage = "logic"

var (
	//go:embed handler.tpl
	handlerTemplate string
	//go:embed sse_handler.tpl
	sseHandlerTemplate string
)

func genHandler(dir, rootPkg, projectPkg string, cfg *config.Config, group spec.Group, route spec.Route) error {
	_ = "STUB: not implemented"
	return nil
}

func genHandlers(dir, rootPkg, projectPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func genHandlerImports(group spec.Group, route spec.Route, parentPkg string) string {
	_ = "STUB: not implemented"
	return ""
}

func getHandlerBaseName(route spec.Route) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getHandlerFolderPath(group spec.Group, route spec.Route) string {
	_ = "STUB: not implemented"
	return ""
}

func getHandlerName(route spec.Route) string { _ = "STUB: not implemented"; return "" }

func getLogicName(route spec.Route) string { _ = "STUB: not implemented"; return "" }
