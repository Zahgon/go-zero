package generate

import (
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

type Context struct {
	Types  []string
	Cache  bool
	Prefix string
	Easy   bool
	Output string
	Cfg    *config.Config
}

func Do(ctx *Context) error { _ = "STUB: not implemented"; return nil }

func generateModel(ctx *Context) error { _ = "STUB: not implemented"; return nil }

func generateCustomModel(ctx *Context) error { _ = "STUB: not implemented"; return nil }

func generateTypes(ctx *Context) error { _ = "STUB: not implemented"; return nil }

func generateError(ctx *Context) error { _ = "STUB: not implemented"; return nil }
