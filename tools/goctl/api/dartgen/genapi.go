package dartgen

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

const apiTemplate = `import 'api.dart';
import '../data/{{with .Info}}{{getBaseName .Title}}{{end}}.dart';
{{with .Service}}
/// {{.Name}}
{{range .Routes}}
/// --{{.Path}}--
///
/// request: {{with .RequestType}}{{.Name}}{{end}}
/// response: {{with .ResponseType}}{{.Name}}{{end}}
Future {{pathToFuncName .Path}}( {{if ne .Method "get"}}{{with .RequestType}}{{.Name}} request,{{end}}{{end}}
    {Function({{with .ResponseType}}{{.Name}}{{end}}) ok,
    Function(String) fail,
    Function eventually}) async {
  await api{{if eq .Method "get"}}Get{{else}}Post{{end}}('{{.Path}}',{{if ne .Method "get"}}request,{{end}}
  	 ok: (data) {
    if (ok != null) ok({{with .ResponseType}}{{.Name}}.fromJson(data){{end}});
  }, fail: fail, eventually: eventually);
}
{{end}}
{{end}}`

const apiTemplateV2 = `import 'api.dart';
import '../data/{{with .Service}}{{.Name}}{{end}}.dart';
{{with .Service}}
/// {{.Name}}
{{range $i, $Route := .Routes}}
/// --{{.Path}}--
///
/// request: {{with .RequestType}}{{.Name}}{{end}}
/// response: {{with .ResponseType}}{{.Name}}{{end}}
Future {{normalizeHandlerName .Handler}}(
	{{if hasUrlPathParams $Route}}{{extractPositionalParamsFromPath $Route}},{{end}}
	{{if ne .Method "get"}}{{with .RequestType}}{{.Name}} request,{{end}}{{end}}
    {Function({{with .ResponseType}}{{.Name}}{{end}})? ok,
    Function(String)? fail,
    Function? eventually}) async {
  await api{{if eq .Method "get"}}Get{{else}}Post{{end}}({{makeDartRequestUrlPath $Route}},{{if ne .Method "get"}}request,{{end}}
  	 ok: (data) {
    if (ok != null) ok({{with .ResponseType}}{{.Name}}.fromJson(data){{end}});
  }, fail: fail, eventually: eventually);
}
{{end}}
{{end}}`

func genApi(dir string, api *spec.ApiSpec, isLegacy bool) error {
	_ = "STUB: not implemented"
	return nil
}

func genApiFile(dir string, isLegacy bool) error { _ = "STUB: not implemented"; return nil }
