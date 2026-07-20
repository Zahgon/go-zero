package gogen

const (
	category                    = "api"
	configTemplateFile          = "config.tpl"
	contextTemplateFile         = "context.tpl"
	etcTemplateFile             = "etc.tpl"
	handlerTemplateFile         = "handler.tpl"
	sseHandlerTemplateFile      = "sse_handler.tpl"
	handlerTestTemplateFile     = "handler_test.tpl"
	logicTemplateFile           = "logic.tpl"
	sseLogicTemplateFile        = "sse_logic.tpl"
	logicTestTemplateFile       = "logic_test.tpl"
	mainTemplateFile            = "main.tpl"
	middlewareImplementCodeFile = "middleware.tpl"
	routesTemplateFile          = "routes.tpl"
	routesAdditionTemplateFile  = "route-addition.tpl"
	typesTemplateFile           = "types.tpl"
	svcTestTemplateFile         = "svc_test.tpl"
	integrationTestTemplateFile = "integration_test.tpl"
)

var templates = map[string]string{
	configTemplateFile:          configTemplate,
	contextTemplateFile:         contextTemplate,
	etcTemplateFile:             etcTemplate,
	handlerTemplateFile:         handlerTemplate,
	sseHandlerTemplateFile:      sseHandlerTemplate,
	handlerTestTemplateFile:     handlerTestTemplate,
	logicTemplateFile:           logicTemplate,
	sseLogicTemplateFile:        sseLogicTemplate,
	logicTestTemplateFile:       logicTestTemplate,
	mainTemplateFile:            mainTemplate,
	middlewareImplementCodeFile: middlewareImplementCode,
	routesTemplateFile:          routesTemplate,
	routesAdditionTemplateFile:  routesAdditionTemplate,
	typesTemplateFile:           typesTemplate,
	svcTestTemplateFile:         svcTestTemplate,
	integrationTestTemplateFile: integrationTestTemplate,
}

func Category() string { _ = "STUB: not implemented"; return "" }

func Clean() error { _ = "STUB: not implemented"; return nil }

func GenTemplates() error { _ = "STUB: not implemented"; return nil }

func RevertTemplate(name string) error { _ = "STUB: not implemented"; return nil }

func Update() error { _ = "STUB: not implemented"; return nil }
