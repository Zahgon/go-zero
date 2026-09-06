package ctx

import (
	"errors"
	"io"

	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
)

const goModuleWithoutGoFiles = "command-line-arguments"

var errInvalidGoMod = errors.New("invalid go module")

type Module struct {
	Path      string
	Main      bool
	Dir       string
	GoMod     string
	GoVersion string
}

func (m *Module) validate() error { _ = "STUB: not implemented"; return nil }

func projectFromGoMod(workDir string) (*ProjectContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getRealModule(workDir string, execRun execx.RunFunc) (*Module, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodePackages(reader io.Reader) ([]Module, error) { _ = "STUB: not implemented"; return nil, nil }
