package gogen

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

const tmpFile = "%s-%d"

var (
	tmpDir = path.Join(os.TempDir(), "goctl")

	VarStringDir string

	VarStringAPI string

	VarStringHome string

	VarStringRemote string

	VarStringBranch string

	VarStringStyle  string
	VarBoolWithTest bool

	VarBoolTypeGroup bool
)

func GoCommand(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func DoGenProject(apiFile, dir, style string, withTest bool) error {
	_ = "STUB: not implemented"
	return nil
}

func DoGenProjectWithModule(apiFile, dir, moduleName, style string, withTest bool) error {
	_ = "STUB: not implemented"
	return nil
}

func backupAndSweep(apiFile string) error { _ = "STUB: not implemented"; return nil }

func sweep() error { _ = "STUB: not implemented"; return nil }
