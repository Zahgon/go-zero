package format

import (
	"bufio"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

const (
	leftParenthesis  = "("
	rightParenthesis = ")"
	leftBrace        = "{"
	rightBrace       = "}"
)

var (
	VarBoolUseStdin bool

	VarBoolSkipCheckDeclare bool

	VarStringDir string

	VarBoolIgnore bool
)

func GoFormatApi(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func apiFormatReader(reader io.Reader, filename string, skipCheckDeclare bool) error {
	_ = "STUB: not implemented"
	return nil
}

func ApiFormatByPath(apiFilePath string, skipCheckDeclare bool) error {
	_ = "STUB: not implemented"
	return nil
}

func apiFormat(data string, skipCheckDeclare bool, filename ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func formatGoTypeDef(line string, scanner *bufio.Scanner, builder *strings.Builder) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func mayInsertStructKeyword(line string, token *int) string { _ = "STUB: not implemented"; return "" }
