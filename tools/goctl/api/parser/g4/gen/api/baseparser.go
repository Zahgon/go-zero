package api

const (
	versionRegex     = `(?m)"v[1-9][0-9]*"`
	importValueRegex = `(?m)"\/?(?:[^/]+\/)*[^/]+.api"`
	tagRegex         = `(?m)\x60[a-z]+:".+"\x60`
)

var (
	holder = struct{}{}
	kind   = map[string]struct{}{
		"bool":       holder,
		"int":        holder,
		"int8":       holder,
		"int16":      holder,
		"int32":      holder,
		"int64":      holder,
		"uint":       holder,
		"uint8":      holder,
		"uint16":     holder,
		"uint32":     holder,
		"uint64":     holder,
		"uintptr":    holder,
		"float32":    holder,
		"float64":    holder,
		"complex64":  holder,
		"complex128": holder,
		"string":     holder,
		"byte":       holder,
		"rune":       holder,
	}
)

func match(p *ApiParserParser, text string) { _ = "STUB: not implemented"; return }

func checkVersion(p *ApiParserParser) { _ = "STUB: not implemented"; return }

func checkImportValue(p *ApiParserParser) { _ = "STUB: not implemented"; return }

func checkKeyValue(p *ApiParserParser) { _ = "STUB: not implemented"; return }

func checkHTTPMethod(p *ApiParserParser) { _ = "STUB: not implemented"; return }

func checkKeyword(p *ApiParserParser) { _ = "STUB: not implemented"; return }

func checkKey(p *ApiParserParser) { _ = "STUB: not implemented"; return }

func IsBasicType(text string) bool { _ = "STUB: not implemented"; return false }

func IsGolangKeyWord(text string, excepts ...string) bool { _ = "STUB: not implemented"; return false }

func isNormal(p *ApiParserParser) bool { _ = "STUB: not implemented"; return false }

func MatchTag(v string) bool { _ = "STUB: not implemented"; return false }

func isInterface(p *ApiParserParser) { _ = "STUB: not implemented"; return }

func getCurrentTokenText(p *ApiParserParser) string { _ = "STUB: not implemented"; return "" }

func setCurrentTokenText(p *ApiParserParser, text string) { _ = "STUB: not implemented"; return }

func notifyErrorListeners(p *ApiParserParser, msg string) { _ = "STUB: not implemented"; return }

func matchRegex(text, str string) bool { _ = "STUB: not implemented"; return false }

func expecting(expecting, found string) string { _ = "STUB: not implemented"; return "" }

func mismatched(expecting, found string) string { _ = "STUB: not implemented"; return "" }
