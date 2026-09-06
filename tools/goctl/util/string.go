package util

var goKeyword = map[string]string{
	"var":         "variable",
	"const":       "constant",
	"package":     "pkg",
	"func":        "function",
	"return":      "rtn",
	"defer":       "dfr",
	"go":          "goo",
	"select":      "slt",
	"struct":      "structure",
	"interface":   "itf",
	"chan":        "channel",
	"type":        "tp",
	"map":         "mp",
	"range":       "rg",
	"break":       "brk",
	"case":        "caz",
	"continue":    "ctn",
	"for":         "fr",
	"fallthrough": "fth",
	"else":        "es",
	"if":          "ef",
	"switch":      "swt",
	"goto":        "gt",
	"default":     "dft",
}

func Title(s string) string { _ = "STUB: not implemented"; return "" }

func Untitle(s string) string { _ = "STUB: not implemented"; return "" }

func Index(slice []string, item string) int { _ = "STUB: not implemented"; return 0 }

func SafeString(in string) string { _ = "STUB: not implemented"; return "" }

func isSafeRune(r rune) bool { _ = "STUB: not implemented"; return false }

func isLetter(r rune) bool { _ = "STUB: not implemented"; return false }

func isNumber(r rune) bool { _ = "STUB: not implemented"; return false }

func EscapeGolangKeyword(s string) string { _ = "STUB: not implemented"; return "" }

func isGolangKeyword(s string) bool { _ = "STUB: not implemented"; return false }

func TrimWhiteSpace(s string) string { _ = "STUB: not implemented"; return "" }

func IsEmptyStringOrWhiteSpace(s string) bool { _ = "STUB: not implemented"; return false }

func FieldsAndTrimSpace(s string, f func(r rune) bool) []string {
	_ = "STUB: not implemented"
	return nil
}
