package docker

import (
	_ "embed"
)

const (
	category           = "docker"
	dockerTemplateFile = "docker.tpl"
)

//go:embed docker.tpl
var dockerTemplate string

func Clean() error { _ = "STUB: not implemented"; return nil }

func GenTemplates() error { _ = "STUB: not implemented"; return nil }

func Category() string { _ = "STUB: not implemented"; return "" }

func RevertTemplate(name string) error { _ = "STUB: not implemented"; return nil }

func Update() error { _ = "STUB: not implemented"; return nil }

func initTemplate() error { _ = "STUB: not implemented"; return nil }
