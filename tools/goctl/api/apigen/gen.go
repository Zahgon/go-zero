package apigen

import (
	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed api.tpl
var apiTemplate string

var (
	VarStringOutput string

	VarStringHome string

	VarStringRemote string

	VarStringBranch string
)

func CreateApiTemplate(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }
