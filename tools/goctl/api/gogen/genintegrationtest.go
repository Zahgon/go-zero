package gogen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

//go:embed integration_test.tpl
var integrationTestTemplate string

func genIntegrationTest(dir, rootPkg, projectPkg string, cfg *config.Config, api *spec.ApiSpec) error {
	_ = "STUB: not implemented"
	return nil
}
