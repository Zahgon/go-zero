package quickstart

import (
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
)

const baseDir = "greet"

var (
	log        = console.NewColorConsole(true)
	projectDir string
)

func cleanWorkSpace(projectDir string) { _ = "STUB: not implemented"; return }

func initProject() { _ = "STUB: not implemented"; return }

func run(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }
