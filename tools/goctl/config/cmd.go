package config

import (
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/internal/cobrax"
)

var (
	Cmd = cobrax.NewCommand("config")

	initCmd  = cobrax.NewCommand("init", cobrax.WithRunE(runConfigInit))
	cleanCmd = cobrax.NewCommand("clean", cobrax.WithRunE(runConfigClean))
)

func init() {
	Cmd.AddCommand(initCmd, cleanCmd)
}

func runConfigInit(*cobra.Command, []string) error { _ = "STUB: not implemented"; return nil }

func runConfigClean(*cobra.Command, []string) error { _ = "STUB: not implemented"; return nil }
