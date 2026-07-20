package new

import (
	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed api.tpl
var apiTemplate string

var (
	VarStringHome string

	VarStringRemote string

	VarStringBranch string

	VarStringStyle string

	VarStringModule string
)

func CreateServiceCommand(_ *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
