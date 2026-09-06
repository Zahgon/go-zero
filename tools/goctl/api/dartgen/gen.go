package dartgen

import (
	"github.com/spf13/cobra"
)

var (
	VarStringDir string

	VarStringAPI string

	VarStringLegacy bool

	VarStringHostname string

	VarStringScheme string
)

func DartCommand(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }
