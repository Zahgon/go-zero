package plugin

import (
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

const pluginArg = "_plugin"

type Plugin struct {
	Api         *spec.ApiSpec
	ApiFilePath string
	Style       string
	Dir         string
}

var (
	VarStringPlugin string

	VarStringDir string

	VarStringAPI string

	VarStringStyle string
)

func PluginCommand(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func prepareArgs() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func getCommand(arg string) (string, bool, error) { _ = "STUB: not implemented"; return "", false, nil }

func downloadFile(filepath, url string) error { _ = "STUB: not implemented"; return nil }

func NewPlugin() (*Plugin, error) { _ = "STUB: not implemented"; return nil, nil }

func getPluginAndArgs(arg string) (string, string) { _ = "STUB: not implemented"; return "", "" }

func trimQuote(in string) string { _ = "STUB: not implemented"; return "" }
