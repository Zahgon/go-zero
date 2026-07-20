package migrate

import (
	"errors"
)

const (
	deprecatedGoZeroMod = "github.com/tal-tech/go-zero"
	deprecatedBuilderx  = "github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
	replacementBuilderx = "github.com/zeromicro/go-zero/core/stores/builder"
	goZeroMod           = "github.com/zeromicro/go-zero"
)

var errInvalidGoMod = errors.New("it's only working for go module")

func editMod(version string, verbose bool) error { _ = "STUB: not implemented"; return nil }

func addRequire(mod string, verbose bool) error { _ = "STUB: not implemented"; return nil }

func removeRequire(mod string, verbose bool) error { _ = "STUB: not implemented"; return nil }

func tidy(verbose bool) error { _ = "STUB: not implemented"; return nil }
