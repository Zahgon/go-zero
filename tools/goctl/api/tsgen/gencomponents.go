package tsgen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

//go:embed components.tpl
var componentsTemplate string

func BuildTypes(types []spec.Type) (string, error) { _ = "STUB: not implemented"; return "", nil }

func genComponents(dir string, api *spec.ApiSpec) error { _ = "STUB: not implemented"; return nil }
