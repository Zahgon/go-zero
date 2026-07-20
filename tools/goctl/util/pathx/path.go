package pathx

const (
	pkgSep           = "/"
	goModeIdentifier = "go.mod"
)

func JoinPackages(pkgs ...string) string { _ = "STUB: not implemented"; return "" }

func MkdirIfNotExist(dir string) error { _ = "STUB: not implemented"; return nil }

func PathFromGoSrc() (string, error) { _ = "STUB: not implemented"; return "", nil }

func FindGoModPath(dir string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func FindProjectPath(loc string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func isLink(name string) (bool, error) { _ = "STUB: not implemented"; return false, nil }
