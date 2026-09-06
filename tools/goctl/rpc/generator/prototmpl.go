package generator

import (
	_ "embed"
)

//go:embed rpc.tpl
var rpcTemplateText string

func ProtoTmpl(out string) error { _ = "STUB: not implemented"; return nil }
