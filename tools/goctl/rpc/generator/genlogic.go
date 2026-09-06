package generator

import (
	_ "embed"

	"github.com/zeromicro/go-zero/core/collection"
	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
)

const logicFunctionTemplate = `{{if .hasComment}}{{.comment}}{{end}}
func (l *{{.logicName}}) {{.method}} ({{if .hasReq}}in {{.request}}{{if .stream}},stream {{.streamBody}}{{end}}{{else}}stream {{.streamBody}}{{end}}) ({{if .hasReply}}{{.response}},{{end}} error) {
	// todo: add your logic here and delete this line
	
	return {{if .hasReply}}&{{.responseType}}{},{{end}} nil
}
`

//go:embed logic.tpl
var logicTemplate string

func (g *Generator) GenLogic(ctx DirContext, proto parser.Proto, cfg *conf.Config,
	c *ZRpcContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) genLogicInCompatibility(ctx DirContext, proto parser.Proto,
	cfg *conf.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) genLogicGroup(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) genLogicFunction(serviceName, goPackage, mainGoPackage, logicName string,
	rpc *parser.RPC, pkgMap map[string]parser.ImportedProto) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func addLogicImports(imports *collection.Set[string], pbImportPath, goPackage, mainGoPackage string,
	rpc *parser.RPC, pkgMap map[string]parser.ImportedProto) {
	_ = "STUB: not implemented"
	return
}
