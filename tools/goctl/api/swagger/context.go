package swagger

import (
	"testing"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

type Context struct {
	UseDefinitions         bool
	WrapCodeMsg            bool
	BizCodeEnumDescription string
}

func testingContext(_ *testing.T) Context { _ = "STUB: not implemented"; return *new(Context) }

func contextFromApi(info spec.Info) Context { _ = "STUB: not implemented"; return *new(Context) }
