package stringx

type node struct {
	children map[rune]*node
	fail     *node
	depth    int
	end      bool
}

func (n *node) add(word string) { _ = "STUB: not implemented"; return }

func (n *node) build() { _ = "STUB: not implemented"; return }

func (n *node) find(chars []rune) []scope { _ = "STUB: not implemented"; return nil }
