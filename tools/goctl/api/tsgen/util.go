package tsgen

import (
	"io"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

const (
	formTagKey   = "form"
	pathTagKey   = "path"
	headerTagKey = "header"
)

func writeProperty(writer io.Writer, member spec.Member, indent int) error {
	_ = "STUB: not implemented"
	return nil
}

func writeIndent(writer io.Writer, indent int) { _ = "STUB: not implemented"; return }

func genTsType(m spec.Member, indent int) (ty string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func goTypeToTs(tp spec.Type, fromPacket bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func addPrefix(tp spec.Type, fromPacket bool) string { _ = "STUB: not implemented"; return "" }

func primitiveType(tp string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func writeType(writer io.Writer, tp spec.Type) error { _ = "STUB: not implemented"; return nil }

func genParamsTypesIfNeed(writer io.Writer, tp spec.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func writeMembers(writer io.Writer, tp spec.Type, isParam bool, indent int) error {
	_ = "STUB: not implemented"
	return nil
}

func writeTagMembers(writer io.Writer, tp spec.Type, tagKey string) error {
	_ = "STUB: not implemented"
	return nil
}
