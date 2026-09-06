package gogen

import (
	_ "embed"
	"io"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

const typesFile = "types"

//go:embed types.tpl
var typesTemplate string

func BuildTypes(types []spec.Type) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getTypeName(tp spec.Type) string { _ = "STUB: not implemented"; return "" }

func genTypesWithGroup(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func writeTypes(dir, baseFilename string, cfg *config.Config, types []spec.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func genTypes(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func writeType(writer io.Writer, tp spec.Type) error { _ = "STUB: not implemented"; return nil }

func writeMember(writer io.Writer, members []spec.Member) error {
	_ = "STUB: not implemented"
	return nil
}
