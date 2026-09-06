package generator

import (
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
)

type ZRpcContext struct {
	Src string

	ProtocCmd string

	ProtoGenGrpcDir string

	ProtoGenGoDir string

	IsGooglePlugin bool

	GoOutput string

	GrpcOutput string

	Output string

	Multiple bool

	IsGenClient bool

	Module string

	NameFromFilename bool

	ProtoPaths []string
}

func (g *Generator) Generate(zctx *ZRpcContext) error { _ = "STUB: not implemented"; return nil }

func resolveImportedProtos(zctx *ZRpcContext) ([]parser.ImportedProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
