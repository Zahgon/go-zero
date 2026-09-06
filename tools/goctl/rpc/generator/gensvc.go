package generator

import (
	_ "embed"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
)

//go:embed svc.tpl
var svcTemplate string

func (g *Generator) GenSvc(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	_ = "STUB: not implemented"
	return nil
}
