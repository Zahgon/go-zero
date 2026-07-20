package stringx

const defaultMask = '*'

type (
	TrieOption func(trie *trieNode)

	Trie interface {
		Filter(text string) (string, []string, bool)
		FindKeywords(text string) []string
	}

	trieNode struct {
		node
		mask rune
	}

	scope struct {
		start int
		stop  int
	}
)

func NewTrie(words []string, opts ...TrieOption) Trie { _ = "STUB: not implemented"; return *new(Trie) }

func (n *trieNode) Filter(text string) (sentence string, keywords []string, found bool) {
	_ = "STUB: not implemented"
	return "", nil, false
}

func (n *trieNode) FindKeywords(text string) []string { _ = "STUB: not implemented"; return nil }

func (n *trieNode) collectKeywords(chars []rune, scopes []scope) []string {
	_ = "STUB: not implemented"
	return nil
}

func (n *trieNode) replaceWithAsterisk(chars []rune, start, stop int) {
	_ = "STUB: not implemented"
	return
}

func WithMask(mask rune) TrieOption { _ = "STUB: not implemented"; return *new(TrieOption) }
