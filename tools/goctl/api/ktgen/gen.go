package ktgen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

var (
	//go:embed apibase.tpl
	apiBaseTemplate string
	//go:embed api.tpl
	apiTemplate string
)

func genBase(dir, pkg string, api *spec.ApiSpec) error { _ = "STUB: not implemented"; return nil }

func genApi(dir, pkg string, api *spec.ApiSpec) error { _ = "STUB: not implemented"; return nil }
