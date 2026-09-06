package apigen

const (
	category        = "api"
	apiTemplateFile = "template.tpl"
)

var templates = map[string]string{
	apiTemplateFile: apiTemplate,
}

func Category() string { _ = "STUB: not implemented"; return "" }

func Clean() error { _ = "STUB: not implemented"; return nil }

func GenTemplates() error { _ = "STUB: not implemented"; return nil }

func RevertTemplate(name string) error { _ = "STUB: not implemented"; return nil }

func Update() error { _ = "STUB: not implemented"; return nil }
