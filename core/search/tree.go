package search

import (
	"errors"
)

const (
	colon = ':'
	slash = '/'
)

var (
	errDupItem = errors.New("duplicated item")

	errDupSlash = errors.New("duplicated slash")

	errEmptyItem = errors.New("empty item")

	errInvalidState = errors.New("search tree is in an invalid state")

	errNotFromRoot = errors.New("path should start with /")

	NotFound Result
)

type (
	innerResult struct {
		key   string
		value string
		named bool
		found bool
	}

	node struct {
		item     any
		children [2]map[string]*node
	}

	Tree struct {
		root *node
	}

	Result struct {
		Item   any
		Params map[string]string
	}
)

func NewTree() *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) Add(route string, item any) error { _ = "STUB: not implemented"; return nil }

func (t *Tree) Search(route string) (Result, bool) {
	_ = "STUB: not implemented"
	return *new(Result), false
}

func (t *Tree) next(n *node, route string, result *Result) bool {
	_ = "STUB: not implemented"
	return false
}

func (nd *node) forEach(fn func(string, *node) bool) bool { _ = "STUB: not implemented"; return false }

func (nd *node) getChildren(route string) map[string]*node { _ = "STUB: not implemented"; return nil }

func add(nd *node, route string, item any) error { _ = "STUB: not implemented"; return nil }

func addParam(result *Result, k, v string) { _ = "STUB: not implemented"; return }

func duplicatedItem(item string) error { _ = "STUB: not implemented"; return nil }

func duplicatedSlash(item string) error { _ = "STUB: not implemented"; return nil }

func match(pat, token string) innerResult { _ = "STUB: not implemented"; return *new(innerResult) }

func newNode(item any) *node { _ = "STUB: not implemented"; return nil }
