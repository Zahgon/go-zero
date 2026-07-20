package stringx

const replaceTimes = 2

type (
	Replacer interface {
		Replace(text string) string
	}

	replacer struct {
		*node
		mapping map[string]string
	}
)

func NewReplacer(mapping map[string]string) Replacer {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

func (r *replacer) Replace(text string) string { _ = "STUB: not implemented"; return "" }

func (r *replacer) doReplace(text string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
