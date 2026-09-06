package ktgen

import (
	"text/template"
)

var funcsMap = template.FuncMap{
	"lowCamelCase":    lowCamelCase,
	"routeToFuncName": routeToFuncName,
	"parseType":       parseType,
	"add":             add,
	"upperCase":       upperCase,
}

func lowCamelCase(s string) string { _ = "STUB: not implemented"; return "" }

func routeToFuncName(method, path string) string { _ = "STUB: not implemented"; return "" }

func parseType(t string) string { _ = "STUB: not implemented"; return "" }

func decomposeType(t string) (result []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func add(a, i int) int { _ = "STUB: not implemented"; return 0 }

func upperCase(s string) string { _ = "STUB: not implemented"; return "" }
