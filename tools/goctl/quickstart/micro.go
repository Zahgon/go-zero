package quickstart

import (
	_ "embed"
)

const protoName = "greet.proto"

var (
	//go:embed idl/greet.proto
	protocContent string
	//go:embed idl/rpc.yaml
	rpcEtcContent string
	zrpcWorkDir   string
)

type serviceImpl struct {
	starter func()
}

func (s serviceImpl) Start() { _ = "STUB: not implemented"; return }

func (s serviceImpl) Stop() { _ = "STUB: not implemented"; return }

func initRPCProto() error { _ = "STUB: not implemented"; return nil }

type micro struct{}

func newMicroService() micro { _ = "STUB: not implemented"; return *new(micro) }

func (m micro) mustStartRPCProject() { _ = "STUB: not implemented"; return }

func (m micro) start() { _ = "STUB: not implemented"; return }
