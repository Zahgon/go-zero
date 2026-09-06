package dartgen

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

const dataTemplate = `// --{{with .APISpec.Info}}{{.Title}}{{end}}--
{{ range .APISpec.Types}}
class {{.Name}}{
	{{range .Members}}
	/// {{.Comment}}
	final {{if isNumberType .Type.Name}}num{{else}}{{.Type.Name}}{{end}} {{lowCamelCase .Name}};
	{{end}}
	{{.Name}}({ {{range .Members}}
		this.{{lowCamelCase .Name}},{{end}}
	});
	factory {{.Name}}.fromJson(Map<String,dynamic> m) {
		return {{.Name}}({{range .Members}}
			{{lowCamelCase .Name}}: {{if isDirectType .Type.Name}}m['{{getPropertyFromMember .}}']{{else if isClassListType .Type.Name}}(m['{{getPropertyFromMember .}}'] as List<dynamic>).map((i) => {{getCoreType .Type.Name}}.fromJson(i)){{else}}{{.Type.Name}}.fromJson(m['{{getPropertyFromMember .}}']){{end}},{{end}}
		);
	}
	Map<String,dynamic> toJson() {
		return { {{range .Members}}
			'{{getPropertyFromMember .}}': {{if isDirectType .Type.Name}}{{lowCamelCase .Name}}{{else if isClassListType .Type.Name}}{{lowCamelCase .Name}}.map((i) => i.toJson()){{else}}{{lowCamelCase .Name}}.toJson(){{end}},{{end}}
		};
	}

	{{ range $.InnerClassList}}
	{{.}}
	{{end}}
}
{{end}}
`

const dataTemplateV2 = `// --{{with .APISpec.Info}}{{.Title}}{{end}}--
{{ range .APISpec.Types}}
class {{.Name}} {
	{{range .Members}}
	{{if .Comment}}{{.Comment}}{{end}}
	final {{if isNumberType .Type.Name}}num{{else}}{{.Type.Name}}{{end}} {{lowCamelCase .Name}};
  {{end}}{{.Name}}({{if .Members}}{
	{{range .Members}}  required this.{{lowCamelCase .Name}},
	{{end}}}{{end}});
	factory {{.Name}}.fromJson(Map<String,dynamic> m) {
		return {{.Name}}(
			{{range .Members}}
				{{lowCamelCase .Name}}: {{appendNullCoalescing .}}
					{{if isAtomicType .Type.Name}}
						m['{{getPropertyFromMember .}}'] {{appendDefaultEmptyValue .Type.Name}}
					{{else if isAtomicListType .Type.Name}}
						m['{{getPropertyFromMember .}}']?.cast<{{getCoreType .Type.Name}}>() {{appendDefaultEmptyValue .Type.Name}}
					{{else if isClassListType .Type.Name}}
						((m['{{getPropertyFromMember .}}'] {{appendDefaultEmptyValue .Type.Name}}) as List<dynamic>).map((i) => {{getCoreType .Type.Name}}.fromJson(i)).toList()
					{{else if isMapType .Type.Name}}
						{{if isNumberType .Type.Name}}num{{else}}{{.Type.Name}}{{end}}.from(m['{{getPropertyFromMember .}}'] ?? {})
					{{else}}
						{{.Type.Name}}.fromJson(m['{{getPropertyFromMember .}}']){{end}}
			,{{end}}
		);
	}
	Map<String,dynamic> toJson() {
		return { {{range .Members}}
			'{{getPropertyFromMember .}}': 
				{{if isDirectType .Type.Name}}
					{{lowCamelCase .Name}}
				{{else if isMapType .Type.Name}}
					{{lowCamelCase .Name}}
				{{else if isClassListType .Type.Name}}
					{{lowCamelCase .Name}}{{if isNullableType .Type.Name}}?{{end}}.map((i) => i{{if isListItemsNullable .Type.Name}}?{{end}}.toJson())
				{{else}}
					{{lowCamelCase .Name}}{{if isNullableType .Type.Name}}?{{end}}.toJson()
				{{end}}
			,{{end}}
		};
	}

	{{ range $.InnerClassList}}
	{{.}}
	{{end}}
}
{{end}}`

type DartSpec struct {
	APISpec        *spec.ApiSpec
	InnerClassList []string
}

func genData(dir string, api *spec.ApiSpec, isLegacy bool) error {
	_ = "STUB: not implemented"
	return nil
}

func genTokens(dir string, isLeagcy bool) error { _ = "STUB: not implemented"; return nil }

func convertDataType(api *spec.ApiSpec, isLegacy bool) (error, *DartSpec) {
	_ = "STUB: not implemented"
	return nil, nil
}
