package template

import (
	_ "embed"
	"fmt"

	"github.com/zeromicro/go-zero/tools/goctl/internal/version"
	"github.com/zeromicro/go-zero/tools/goctl/util"
)

//go:embed tpl/customized.tpl
var Customized string

//go:embed tpl/var.tpl
var Vars string

//go:embed tpl/types.tpl
var Types string

//go:embed tpl/tag.tpl
var Tag string

//go:embed tpl/table-name.tpl
var TableName string

//go:embed tpl/model-new.tpl
var New string

//go:embed tpl/model.tpl
var ModelCustom string

var ModelGen = fmt.Sprintf(`%s
// versions:
//  goctl version: %s

package {{.pkg}}
{{.imports}}
{{.vars}}
{{.types}}
{{.new}}
{{.delete}}
{{.find}}
{{.insert}}
{{.update}}
{{.extraMethod}}
{{.tableName}}
{{.customized}}
`, util.DoNotEditHead, version.BuildVersion)

//go:embed tpl/insert.tpl
var Insert string

//go:embed tpl/interface-insert.tpl
var InsertMethod string

//go:embed tpl/update.tpl
var Update string

//go:embed tpl/interface-update.tpl
var UpdateMethod string

//go:embed tpl/import.tpl
var Imports string

//go:embed tpl/import-no-cache.tpl
var ImportsNoCache string

//go:embed tpl/find-one.tpl
var FindOne string

//go:embed tpl/find-one-by-field.tpl
var FindOneByField string

//go:embed tpl/find-one-by-field-extra-method.tpl
var FindOneByFieldExtraMethod string

//go:embed tpl/interface-find-one.tpl
var FindOneMethod string

//go:embed tpl/interface-find-one-by-field.tpl
var FindOneByFieldMethod string

//go:embed tpl/field.tpl
var Field string

//go:embed tpl/err.tpl
var Error string

//go:embed tpl/delete.tpl
var Delete string

//go:embed tpl/interface-delete.tpl
var DeleteMethod string
