package docgen

import (
	"github.com/spf13/cobra"
)

var (
	VarStringDir string

	VarStringOutput string
)

func DocCommand(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func filePathWalkDir(root string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
