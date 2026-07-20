package bug

import (
	"github.com/spf13/cobra"
)

const (
	windows = "windows"
	darwin  = "darwin"

	windowsOpen = "start"
	darwinOpen  = "open"
	linuxOpen   = "xdg-open"

	os           = "OS"
	arch         = "ARCH"
	goctlVersion = "GOCTL_VERSION"
	goVersion    = "GO_VERSION"
)

var openCmd = map[string]string{
	windows: windowsOpen,
	darwin:  darwinOpen,
}

func runE(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }
