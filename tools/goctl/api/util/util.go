package util

import (
	"io"
	"os"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func MaybeCreateFile(dir, subdir, file string) (fp *os.File, created bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func WrapErr(err error, message string) error { _ = "STUB: not implemented"; return nil }

func Copy(src, dst string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func ComponentName(api *spec.ApiSpec) string { _ = "STUB: not implemented"; return "" }

func WriteIndent(writer io.Writer, indent int) { _ = "STUB: not implemented"; return }

func RemoveComment(line string) string { _ = "STUB: not implemented"; return "" }
