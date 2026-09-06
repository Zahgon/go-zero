package generator

import (
	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
)

const (
	wd       = "wd"
	etc      = "etc"
	internal = "internal"
	config   = "config"
	logic    = "logic"
	server   = "server"
	svc      = "svc"
	pb       = "pb"
	protoGo  = "proto-go"
	call     = "call"
)

type (
	DirContext interface {
		GetCall() Dir
		GetEtc() Dir
		GetInternal() Dir
		GetConfig() Dir
		GetLogic() Dir
		GetServer() Dir
		GetSvc() Dir
		GetPb() Dir
		GetProtoGo() Dir
		GetMain() Dir
		GetServiceName() stringx.String
		SetPbDir(pbDir, grpcDir string)
	}

	Dir struct {
		Base            string
		Filename        string
		Package         string
		GetChildPackage func(childPath string) (string, error)
	}

	defaultDirContext struct {
		inner       map[string]Dir
		serviceName stringx.String
		ctx         *ctx.ProjectContext
	}
)

func mkdir(ctx *ctx.ProjectContext, proto parser.Proto, conf *conf.Config, c *ZRpcContext) (DirContext,
	error) {
	_ = "STUB: not implemented"
	return *new(DirContext), nil
}

func (d *defaultDirContext) SetPbDir(pbDir, grpcDir string) { _ = "STUB: not implemented"; return }

func (d *defaultDirContext) GetCall() Dir { _ = "STUB: not implemented"; return *new(Dir) }

func (d *defaultDirContext) GetEtc() Dir { _ = "STUB: not implemented"; return *new(Dir) }

func (d *defaultDirContext) GetInternal() Dir { _ = "STUB: not implemented"; return *new(Dir) }

func (d *defaultDirContext) GetConfig() Dir { _ = "STUB: not implemented"; return *new(Dir) }

func (d *defaultDirContext) GetLogic() Dir { _ = "STUB: not implemented"; return *new(Dir) }

func (d *defaultDirContext) GetServer() Dir { _ = "STUB: not implemented"; return *new(Dir) }

func (d *defaultDirContext) GetSvc() Dir { _ = "STUB: not implemented"; return *new(Dir) }

func (d *defaultDirContext) GetPb() Dir { _ = "STUB: not implemented"; return *new(Dir) }

func (d *defaultDirContext) GetProtoGo() Dir { _ = "STUB: not implemented"; return *new(Dir) }

func (d *defaultDirContext) GetMain() Dir { _ = "STUB: not implemented"; return *new(Dir) }

func (d *defaultDirContext) GetServiceName() stringx.String {
	_ = "STUB: not implemented"
	return *new(stringx.String)
}

func (d *Dir) Valid() bool { _ = "STUB: not implemented"; return false }

func determineServiceName(proto parser.Proto, c *ZRpcContext) string {
	_ = "STUB: not implemented"
	return ""
}
