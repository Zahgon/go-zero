package swagger

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func rangeValueFromOptions(options []string) (minimum *float64, maximum *float64, exclusiveMinimum bool, exclusiveMaximum bool) {
	_ = "STUB: not implemented"
	return nil, nil, false, false
}

func enumsValueFromOptions(options []string) []any { _ = "STUB: not implemented"; return nil }

func defValueFromOptions(ctx Context, options []string, apiType spec.Type) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func exampleValueFromOptions(ctx Context, options []string, apiType spec.Type) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func valueFromOptions(_ Context, options []string, key string, tp string) any {
	_ = "STUB: not implemented"
	return *new(any)
}
