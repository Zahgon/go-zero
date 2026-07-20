package zipx

import (
	"archive/zip"
)

func Unpacking(name, destPath string, mapper func(f *zip.File) bool) error {
	_ = "STUB: not implemented"
	return nil
}

func fileCopy(file *zip.File, destPath string) error { _ = "STUB: not implemented"; return nil }
