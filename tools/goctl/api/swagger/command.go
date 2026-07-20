package swagger

import (
	"github.com/spf13/cobra"
)

var (
	VarStringAPI string

	VarStringDir string

	VarStringFilename string

	VarBoolYaml bool
)

func Command(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }
