package format

import (
	"errors"
)

const (
	flagGo   = "GO"
	flagZero = "ZERO"

	unknown style = iota
	title
	lower
	upper
)

var ErrNamingFormat = errors.New("unsupported format")

type (
	styleFormat struct {
		before    string
		through   string
		after     string
		goStyle   style
		zeroStyle style
	}

	style int
)

func FileNamingFormat(format, content string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func doFormat(f styleFormat, content string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func transferTo(in string, style style) string { _ = "STUB: not implemented"; return "" }

func split(content string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func getStyle(flag string) (style, error) { _ = "STUB: not implemented"; return *new(style), nil }
