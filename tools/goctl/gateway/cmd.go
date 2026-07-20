package gateway

import (
	_ "embed"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/internal/cobrax"
)

var (
	varStringHome   string
	varStringRemote string
	varStringBranch string
	varStringDir    string

	Cmd = cobrax.NewCommand("gateway", cobrax.WithRunE(generateGateway))
)

func init() {
	Cmd.PersistentFlags().StringVar(&varStringHome, "home")
	Cmd.PersistentFlags().StringVar(&varStringRemote, "remote")
	Cmd.PersistentFlags().StringVar(&varStringBranch, "branch")
	Cmd.PersistentFlags().StringVar(&varStringDir, "dir")
}

func generateGateway(*cobra.Command, []string) error { _ = "STUB: not implemented"; return nil }
