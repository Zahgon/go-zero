package generator

import (
	_ "embed"

	"github.com/emicklei/proto"
	"github.com/zeromicro/go-zero/core/collection"
	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
)

const (
	callInterfaceFunctionTemplate = `{{if .hasComment}}{{.comment}}
{{end}}{{.method}}(ctx context.Context{{if .hasReq}}, in *{{.pbRequest}}{{end}}, opts ...grpc.CallOption) ({{if .notStream}}*{{.pbResponse}}, {{else}}{{.streamBody}},{{end}} error)`

	callFunctionTemplate = `
{{if .hasComment}}{{.comment}}{{end}}
func (m *default{{.serviceName}}) {{.method}}(ctx context.Context{{if .hasReq}}, in *{{.pbRequest}}{{end}}, opts ...grpc.CallOption) ({{if .notStream}}*{{.pbResponse}}, {{else}}{{.streamBody}},{{end}} error) {
	client := {{if .isCallPkgSameToGrpcPkg}}{{else}}{{.package}}.{{end}}New{{.rpcServiceName}}Client(m.cli.Conn())
	return client.{{.method}}(ctx{{if .hasReq}}, in{{end}}, opts...)
}
`
)

//go:embed call.tpl
var callTemplateText string

func (g *Generator) GenCall(ctx DirContext, proto parser.Proto, cfg *conf.Config,
	c *ZRpcContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) genCallGroup(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) genCallInCompatibility(ctx DirContext, proto parser.Proto,
	cfg *conf.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func getMessageName(msg proto.Message) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) genFunction(goPackage, mainGoPackage, serviceName string, service parser.Service,
	isCallPkgSameToGrpcPkg bool, pkgMap map[string]parser.ImportedProto,
	alias, extraImports *collection.Set[string]) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) getInterfaceFuncs(goPackage, mainGoPackage string, service parser.Service,
	isCallPkgSameToGrpcPkg bool, pkgMap map[string]parser.ImportedProto,
	extraImports *collection.Set[string]) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectServiceUsedTypes(messages []parser.Message, service parser.Service) *collection.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

func collectMessageDependencies(protoType string, messageByName map[string]*proto.Message,
	usedTypes *collection.Set[string]) {
	_ = "STUB: not implemented"
	return
}

func messageTypeCandidates(protoType string) []string { _ = "STUB: not implemented"; return nil }

func buildExtraImportLines(extraImports *collection.Set[string]) string {
	_ = "STUB: not implemented"
	return ""
}
