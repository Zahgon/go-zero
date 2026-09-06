package quickstart

import (
	_ "embed"
)

var (
	//go:embed idl/greet.api
	apiContent string
	//go:embed idl/svc.tpl
	svcContent string
	//go:embed idl/apilogic.tpl
	apiLogicContent string
	//go:embed idl/api.yaml
	apiEtcContent string

	apiWorkDir string
	rpcWorkDir string
)

func initAPIFlags() error { _ = "STUB: not implemented"; return nil }

type mono struct {
	callRPC bool
}

func newMonoService(callRPC bool) mono { _ = "STUB: not implemented"; return *new(mono) }

func (m mono) createAPIProject() { _ = "STUB: not implemented"; return }

func (m mono) start() { _ = "STUB: not implemented"; return }
