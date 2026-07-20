package tsgen

import (
	"github.com/spf13/cobra"
)

var (
	VarStringDir string

	VarStringAPI string

	VarStringWebAPI string

	VarStringCaller string

	VarBoolUnWrap bool
)

func TsCommand(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }
