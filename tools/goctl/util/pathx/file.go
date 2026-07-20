package pathx

import (
	"os"
)

const (
	NL              = "\n"
	goctlDir        = ".goctl"
	gitDir          = ".git"
	autoCompleteDir = ".auto_complete"
	cacheDir        = "cache"
)

var goctlHome string

func RegisterGoctlHome(home string) { _ = "STUB: not implemented"; return }

func CreateIfNotExist(file string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveIfExist(filename string) error { _ = "STUB: not implemented"; return nil }

func RemoveOrQuit(filename string) error { _ = "STUB: not implemented"; return nil }

func FileExists(file string) bool { _ = "STUB: not implemented"; return false }

func FileNameWithoutExt(file string) string { _ = "STUB: not implemented"; return "" }

func GetGoctlHome() (home string, err error) { _ = "STUB: not implemented"; return "", nil }

func GetDefaultGoctlHome() (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetGitHome() (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetAutoCompleteHome() (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetCacheDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetTemplateDir(category string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func InitTemplates(category string, templates map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateTemplate(category, name, content string) error { _ = "STUB: not implemented"; return nil }

func Clean(category string) error { _ = "STUB: not implemented"; return nil }

func LoadTemplate(category, file, builtin string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func SameFile(path1, path2 string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func createTemplate(file, content string, force bool) error { _ = "STUB: not implemented"; return nil }

func MustTempDir() string { _ = "STUB: not implemented"; return "" }

func Copy(src, dest string) error { _ = "STUB: not implemented"; return nil }

func Hash(file string) (string, error) { _ = "STUB: not implemented"; return "", nil }
