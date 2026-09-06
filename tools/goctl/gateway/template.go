package gateway

import (
	_ "embed"
)

const (
	category            = "gateway"
	etcTemplateFileFile = "etc.tpl"
	mainTemplateFile    = "main.tpl"
)

//go:embed conf.yml
var etcTemplate string

//go:embed gateway.tpl
var mainTemplate string

var templates = map[string]string{
	etcTemplateFileFile: etcTemplate,
	mainTemplateFile:    mainTemplate,
}

func GenTemplates() error { _ = "STUB: not implemented"; return nil }

func RevertTemplate(name string) error { _ = "STUB: not implemented"; return nil }

func Clean() error { _ = "STUB: not implemented"; return nil }

func Update() error { _ = "STUB: not implemented"; return nil }

func Category() string { _ = "STUB: not implemented"; return "" }
