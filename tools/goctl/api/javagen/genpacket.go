package javagen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

//go:embed packet.tpl
var packetTemplate string

func genPacket(dir, packetName string, api *spec.ApiSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func createWith(dir string, api *spec.ApiSpec, route spec.Route, packetName string) error {
	_ = "STUB: not implemented"
	return nil
}

func doc(route spec.Route) string { _ = "STUB: not implemented"; return "" }

func getImports(api *spec.ApiSpec, packetName string) string { _ = "STUB: not implemented"; return "" }

func paramsSet(route spec.Route) string { _ = "STUB: not implemented"; return "" }

func paramsForRoute(route spec.Route) string { _ = "STUB: not implemented"; return "" }

func declarationForRoute(route spec.Route) string { _ = "STUB: not implemented"; return "" }

func processUri(route spec.Route) string { _ = "STUB: not implemented"; return "" }

func formString(route spec.Route) string { _ = "STUB: not implemented"; return "" }
