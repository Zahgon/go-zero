package javagen

import (
	"io"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func writeProperty(writer io.Writer, member spec.Member, indent int) error {
	_ = "STUB: not implemented"
	return nil
}

func writeDefaultValue(writer io.Writer, member spec.Member) error {
	_ = "STUB: not implemented"
	return nil
}

func writeIndent(writer io.Writer, indent int) { _ = "STUB: not implemented"; return }

func indentString(indent int) string { _ = "STUB: not implemented"; return "" }

func specTypeToJava(tp spec.Type) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getBaseType(valueType string) string { _ = "STUB: not implemented"; return "" }

func primitiveType(tp string) (string, bool) { _ = "STUB: not implemented"; return "", false }
