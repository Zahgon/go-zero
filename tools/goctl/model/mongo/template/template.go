package template

import _ "embed"

var (

	//go:embed model.tpl
	ModelText string

	//go:embed model_custom.tpl
	ModelCustomText string

	//go:embed types.tpl
	ModelTypesText string

	//go:embed error.tpl
	Error string
)
