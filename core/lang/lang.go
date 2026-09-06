package lang

import (
	"reflect"
)

var Placeholder PlaceholderType

type (
	AnyType = any

	PlaceholderType = struct{}
)

func Repr(v any) string { _ = "STUB: not implemented"; return "" }

func reprOfValue(val reflect.Value) string { _ = "STUB: not implemented"; return "" }
