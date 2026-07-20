package parser

import (
	"github.com/emicklei/proto"
)

type (
	Services []Service

	Service struct {
		*proto.Service
		RPC []*RPC
	}
)

func (s Services) validate(filename string, multipleOpt ...bool) error {
	_ = "STUB: not implemented"
	return nil
}
