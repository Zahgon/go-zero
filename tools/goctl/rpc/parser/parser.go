package parser

import (
	"errors"
)

type (
	DefaultProtoParser struct{}
)

var ErrGoPackage = errors.New(`option go_package = "" field is not filled in`)

func NewDefaultProtoParser() *DefaultProtoParser { _ = "STUB: not implemented"; return nil }

func (p *DefaultProtoParser) Parse(src string, multiple ...bool) (Proto, error) {
	_ = "STUB: not implemented"
	return *new(Proto), nil
}

func GoSanitized(s string) string { _ = "STUB: not implemented"; return "" }

func CamelCase(s string) string { _ = "STUB: not implemented"; return "" }

func isASCIILower(c byte) bool { _ = "STUB: not implemented"; return false }

func isASCIIDigit(c byte) bool { _ = "STUB: not implemented"; return false }
