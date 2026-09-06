package javagen

import (
	_ "embed"
	"io"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

const (
	httpResponseData = "import com.xhb.core.response.HttpResponseData;"
	httpData         = "import com.xhb.core.packet.HttpData;"
)

var (
	//go:embed component.tpl
	componentTemplate string
	//go:embed getset.tpl
	getSetTemplate string
	//go:embed bool.tpl
	boolTemplate string
)

type componentsContext struct {
	api           *spec.ApiSpec
	requestTypes  []spec.Type
	responseTypes []spec.Type
	imports       []string
	members       []spec.Member
}

func genComponents(dir, packetName string, api *spec.ApiSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *componentsContext) createComponent(dir, packetName string, ty spec.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *componentsContext) checkStruct(ty spec.Type) (spec.DefineStruct, bool, error) {
	_ = "STUB: not implemented"
	return *new(spec.DefineStruct), false, nil
}

func (c *componentsContext) buildProperties(defineStruct spec.DefineStruct) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *componentsContext) buildGetterSetter(defineStruct spec.DefineStruct) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *componentsContext) writeType(writer io.Writer, defineStruct spec.DefineStruct) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *componentsContext) writeMembers(writer io.Writer, tp spec.Type, indent int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *componentsContext) buildConstructor() (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (c *componentsContext) genGetSet(writer io.Writer, indent int) error {
	_ = "STUB: not implemented"
	return nil
}

func formatSource(source string) string { _ = "STUB: not implemented"; return "" }
