package mongo

import (
	"github.com/spf13/cobra"
)

var (
	VarStringSliceType []string

	VarStringDir string

	VarBoolCache bool

	VarStringPrefix string

	VarBoolEasy bool

	VarStringStyle string

	VarStringHome string

	VarStringRemote string

	VarStringBranch string
)

func Action(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }
