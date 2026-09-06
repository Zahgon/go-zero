package gogen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

//go:embed main.tpl
var mainTemplate string

func genMain(dir, rootPkg, projectPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func genMainImports(parentPkg string) string { _ = "STUB: not implemented"; return "" }
