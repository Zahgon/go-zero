package gogen

import (
	"io"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

type fileGenConfig struct {
	dir             string
	subdir          string
	filename        string
	templateName    string
	category        string
	templateFile    string
	builtinTemplate string
	data            any
}

func genFile(c fileGenConfig) error { _ = "STUB: not implemented"; return nil }

func writeProperty(writer io.Writer, name, tag, comment string, tp spec.Type, indent int) error {
	_ = "STUB: not implemented"
	return nil
}

func getAuths(api *spec.ApiSpec) []string { _ = "STUB: not implemented"; return nil }

func getJwtTrans(api *spec.ApiSpec) []string { _ = "STUB: not implemented"; return nil }

func getMiddleware(api *spec.ApiSpec) []string { _ = "STUB: not implemented"; return nil }

func responseGoTypeName(r spec.Route, pkg ...string) string { _ = "STUB: not implemented"; return "" }

func requestGoTypeName(r spec.Route, pkg ...string) string { _ = "STUB: not implemented"; return "" }

func golangExpr(ty spec.Type, pkg ...string) string { _ = "STUB: not implemented"; return "" }

func getDoc(doc string) string { _ = "STUB: not implemented"; return "" }
