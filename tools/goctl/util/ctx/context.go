package ctx

import (
	"errors"
)

var errModuleCheck = errors.New("the work directory must be found in the go mod or the $GOPATH")

type ProjectContext struct {
	WorkDir string

	Name string

	Path string

	Dir string
}

func Prepare(workDir string) (*ProjectContext, error) { _ = "STUB: not implemented"; return nil, nil }

func PrepareWithModule(workDir string, moduleName string) (*ProjectContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func background(workDir string) (*ProjectContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
