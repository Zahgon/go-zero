package stringx

import (
	"errors"
)

var (
	ErrInvalidStartPosition = errors.New("start position is invalid")

	ErrInvalidStopPosition = errors.New("stop position is invalid")
)

func Contains(list []string, str string) bool { _ = "STUB: not implemented"; return false }

func Filter(s string, remove func(r rune) bool) string { _ = "STUB: not implemented"; return "" }

func FirstN(s string, n int, ellipsis ...string) string { _ = "STUB: not implemented"; return "" }

func HasEmpty(args ...string) bool { _ = "STUB: not implemented"; return false }

func Join(sep byte, elem ...string) string { _ = "STUB: not implemented"; return "" }

func NotEmpty(args ...string) bool { _ = "STUB: not implemented"; return false }

func Remove(strings []string, strs ...string) []string { _ = "STUB: not implemented"; return nil }

func Reverse(s string) string { _ = "STUB: not implemented"; return "" }

func Substr(str string, start, stop int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func TakeOne(valid, or string) string { _ = "STUB: not implemented"; return "" }

func TakeWithPriority(fns ...func() string) string { _ = "STUB: not implemented"; return "" }

func ToCamelCase(s string) string { _ = "STUB: not implemented"; return "" }

func Union(first, second []string) []string { _ = "STUB: not implemented"; return nil }
