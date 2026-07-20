package generator

import (
	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
)

type Generator struct {
	log     console.Console
	cfg     *conf.Config
	verbose bool
}

func NewGenerator(style string, verbose bool) *Generator { _ = "STUB: not implemented"; return nil }

func (g *Generator) Prepare() error { _ = "STUB: not implemented"; return nil }
