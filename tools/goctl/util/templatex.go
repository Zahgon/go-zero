package util

import (
	"bytes"
	"text/template"
)

const regularPerm = 0o666

type DefaultTemplate struct {
	name    string
	text    string
	goFmt   bool
	funcMap template.FuncMap
}

func With(name string) *DefaultTemplate { _ = "STUB: not implemented"; return nil }

func (t *DefaultTemplate) Parse(text string) *DefaultTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultTemplate) GoFmt(format bool) *DefaultTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultTemplate) SaveTo(data any, path string, forceUpdate bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultTemplate) Execute(data any) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *DefaultTemplate) AddFunc(funcName string, function any) *DefaultTemplate {
	_ = "STUB: not implemented"
	return nil
}

func IsTemplateVariable(text string) bool { _ = "STUB: not implemented"; return false }

func TemplateVariable(text string) string { _ = "STUB: not implemented"; return "" }
