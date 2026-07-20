package generator

import (
	_ "embed"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
)

//go:embed main.tpl
var mainTemplate string

type MainServiceTemplateData struct {
	GRPCService string
	Service     string
	ServerPkg   string
	Pkg         string
}

func (g *Generator) GenMain(ctx DirContext, proto parser.Proto, cfg *conf.Config,
	c *ZRpcContext) error {
	_ = "STUB: not implemented"
	return nil
}
