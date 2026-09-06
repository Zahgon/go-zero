package stringx

var WhiteSpace = []rune{'\n', '\t', '\f', '\v', ' '}

type String struct {
	source string
}

func From(data string) String { _ = "STUB: not implemented"; return *new(String) }

func (s String) IsEmptyOrSpace() bool { _ = "STUB: not implemented"; return false }

func (s String) Lower() string { _ = "STUB: not implemented"; return "" }

func (s String) Upper() string { _ = "STUB: not implemented"; return "" }

func (s String) ReplaceAll(old, new string) string { _ = "STUB: not implemented"; return "" }

func (s String) Source() string { _ = "STUB: not implemented"; return "" }

func (s String) Title() string { _ = "STUB: not implemented"; return "" }

func (s String) ToCamel() string { _ = "STUB: not implemented"; return "" }

func (s String) ToSnake() string { _ = "STUB: not implemented"; return "" }

func (s String) Untitle() string { _ = "STUB: not implemented"; return "" }

func (s String) splitBy(fn func(r rune) bool, remove bool) []string {
	_ = "STUB: not implemented"
	return nil
}

func ContainsAny(s string, runes ...rune) bool { _ = "STUB: not implemented"; return false }

func ContainsWhiteSpace(s string) bool { _ = "STUB: not implemented"; return false }

func IsWhiteSpace(text string) bool { _ = "STUB: not implemented"; return false }
