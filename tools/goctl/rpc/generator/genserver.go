package generator

import (
	_ "embed"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
)

const functionTemplate = `
{{if .hasComment}}{{.comment}}{{end}}
func (s *{{.server}}Server) {{.method}} ({{if .notStream}}ctx context.Context,{{if .hasReq}} in {{.request}}{{end}}{{else}}{{if .hasReq}} in {{.request}},{{end}}stream {{.streamBody}}{{end}}) ({{if .notStream}}{{.response}},{{end}}error) {
	l := {{.logicPkg}}.New{{.logicName}}({{if .notStream}}ctx,{{else}}stream.Context(),{{end}}s.svcCtx)
	return l.{{.method}}({{if .hasReq}}in{{if .stream}} ,stream{{end}}{{else}}{{if .stream}}stream{{end}}{{end}})
}
`

//go:embed server.tpl
var serverTemplate string

func (g *Generator) GenServer(ctx DirContext, proto parser.Proto, cfg *conf.Config,
	c *ZRpcContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) genServerGroup(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) genServerInCompatibility(ctx DirContext, proto parser.Proto,
	cfg *conf.Config, c *ZRpcContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) genFunctions(goPackage, mainGoPackage string, service parser.Service,
	multiple bool, pkgMap map[string]parser.ImportedProto) ([]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
