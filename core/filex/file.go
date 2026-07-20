package filex

import (
	"os"
)

const bufSize = 1024

func FirstLine(filename string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func LastLine(filename string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func firstLine(file *os.File) (string, error) { _ = "STUB: not implemented"; return "", nil }

func lastLine(filename string, file *os.File) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
