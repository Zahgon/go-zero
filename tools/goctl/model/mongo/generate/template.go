package generate

import (
	"github.com/zeromicro/go-zero/tools/goctl/model/mongo/template"
)

const (
	category                = "mongo"
	modelTemplateFile       = "model.tpl"
	modelCustomTemplateFile = "model_custom.tpl"
	modelTypesTemplateFile  = "model_types.tpl"
	errTemplateFile         = "err.tpl"
)

var templates = map[string]string{
	modelTemplateFile:       template.ModelText,
	modelCustomTemplateFile: template.ModelCustomText,
	modelTypesTemplateFile:  template.ModelTypesText,
	errTemplateFile:         template.Error,
}

func Category() string { _ = "STUB: not implemented"; return "" }

func Clean() error { _ = "STUB: not implemented"; return nil }

func Templates() error { _ = "STUB: not implemented"; return nil }

func RevertTemplate(name string) error { _ = "STUB: not implemented"; return nil }

func Update() error { _ = "STUB: not implemented"; return nil }
