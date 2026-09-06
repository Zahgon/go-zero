package name

type NamingStyle = string

const (
	NamingLower NamingStyle = "lower"

	NamingCamel NamingStyle = "camel"

	NamingSnake NamingStyle = "snake"
)

func IsNamingValid(namingStyle string) (NamingStyle, bool) {
	_ = "STUB: not implemented"
	return *new(NamingStyle), false
}

func FormatFilename(filename string, style NamingStyle) string {
	_ = "STUB: not implemented"
	return ""
}
