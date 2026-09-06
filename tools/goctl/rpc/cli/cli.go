package cli

import (
	"github.com/spf13/cobra"
)

var (
	VarStringOutput string

	VarStringHome string

	VarStringRemote string

	VarStringBranch string

	VarStringSliceGoOut []string

	VarStringSliceGoGRPCOut []string

	VarStringSlicePlugin []string

	VarStringSliceProtoPath []string

	VarStringSliceGoOpt []string

	VarStringSliceGoGRPCOpt []string

	VarStringStyle string

	VarStringZRPCOut string

	VarBoolIdea bool

	VarBoolVerbose bool

	VarBoolMultiple bool

	VarBoolClient bool

	VarStringModule string

	VarBoolNameFromFilename bool
)

func RPCNew(_ *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func RPCTemplate(latest bool) error { _ = "STUB: not implemented"; return nil }
