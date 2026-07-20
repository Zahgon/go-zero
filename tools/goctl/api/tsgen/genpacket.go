package tsgen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

//go:embed handler.tpl
var handlerTemplate string

func genHandler(dir, webAPI, caller string, api *spec.ApiSpec, unwrapAPI bool) error {
	_ = "STUB: not implemented"
	return nil
}

func genAPI(api *spec.ApiSpec, caller string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func paramsForRoute(route spec.Route) string { _ = "STUB: not implemented"; return "" }

func commentForRoute(route spec.Route) string { _ = "STUB: not implemented"; return "" }

func callParamsForRoute(route spec.Route, group spec.Group) string {
	_ = "STUB: not implemented"
	return ""
}

func pathForRoute(route spec.Route, group spec.Group) string { _ = "STUB: not implemented"; return "" }

func pathHasParams(route spec.Route) bool { _ = "STUB: not implemented"; return false }

func hasRequestBody(route spec.Route) bool { _ = "STUB: not implemented"; return false }

func hasRequestPath(route spec.Route) bool { _ = "STUB: not implemented"; return false }

func hasRequestHeader(route spec.Route) bool { _ = "STUB: not implemented"; return false }
