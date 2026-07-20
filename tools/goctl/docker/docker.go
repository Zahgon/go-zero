package docker

import (
	"github.com/spf13/cobra"
)

const (
	dockerfileName = "Dockerfile"
	etcDir         = "etc"
	yamlEtx        = ".yaml"
)

type Docker struct {
	Chinese     bool
	GoMainFrom  string
	GoRelPath   string
	GoFile      string
	ExeFile     string
	BaseImage   string
	HasPort     bool
	Port        int
	Argument    string
	Version     string
	HasTimezone bool
	Timezone    string
}

func dockerCommand(_ *cobra.Command, _ []string) (err error) { _ = "STUB: not implemented"; return nil }

func findConfig(file, dir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func generateDockerfile(goFile, base string, port int, version, timezone string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func getFilePath(file string) (string, error) { _ = "STUB: not implemented"; return "", nil }
