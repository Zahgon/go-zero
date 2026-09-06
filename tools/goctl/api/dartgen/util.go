package dartgen

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

const (
	formTagKey   = "form"
	pathTagKey   = "path"
	headerTagKey = "header"
)

func normalizeHandlerName(handlerName string) string { _ = "STUB: not implemented"; return "" }

func lowCamelCase(s string) string { _ = "STUB: not implemented"; return "" }

func getBaseName(str string) string { _ = "STUB: not implemented"; return "" }

func getPropertyFromMember(member spec.Member) string { _ = "STUB: not implemented"; return "" }

func isDirectType(s string) bool { _ = "STUB: not implemented"; return false }

func isAtomicType(s string) bool { _ = "STUB: not implemented"; return false }

func isNumberType(s string) bool { _ = "STUB: not implemented"; return false }

func isListType(s string) bool { _ = "STUB: not implemented"; return false }

func isClassListType(s string) bool { _ = "STUB: not implemented"; return false }

func isAtomicListType(s string) bool { _ = "STUB: not implemented"; return false }

func isListItemsNullable(s string) bool { _ = "STUB: not implemented"; return false }

func isMapType(s string) bool { _ = "STUB: not implemented"; return false }

func isNullableType(s string) bool { _ = "STUB: not implemented"; return false }

func appendNullCoalescing(member spec.Member) string { _ = "STUB: not implemented"; return "" }

func appendDefaultEmptyValue(s string) string { _ = "STUB: not implemented"; return "" }

func getCoreType(s string) string { _ = "STUB: not implemented"; return "" }

func fileExists(path string) bool { _ = "STUB: not implemented"; return false }

func buildSpecType(tp spec.Type, name string) spec.Type {
	_ = "STUB: not implemented"
	return *new(spec.Type)
}

func specTypeToDart(tp spec.Type) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getBaseType(valueType string) string { _ = "STUB: not implemented"; return "" }

func primitiveType(tp string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func hasUrlPathParams(route spec.Route) bool { _ = "STUB: not implemented"; return false }

func extractPositionalParamsFromPath(route spec.Route) string { _ = "STUB: not implemented"; return "" }

func makeDartRequestUrlPath(route spec.Route) string { _ = "STUB: not implemented"; return "" }
