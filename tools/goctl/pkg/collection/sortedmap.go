package sortedmap

import (
	"container/list"
	"errors"
)

var (
	ErrInvalidKVExpression = errors.New(`invalid key-value expression`)
	ErrInvalidKVS          = errors.New("the length of kv must be an even number")
)

type KV []any

type SortedMap struct {
	kv   *list.List
	keys map[any]*list.Element
}

func New() *SortedMap { _ = "STUB: not implemented"; return nil }

func (m *SortedMap) SetExpression(expression string) (key, value any, err error) {
	_ = "STUB: not implemented"
	return *new(any), *new(any), nil
}

func (m *SortedMap) SetKV(key, value any) { _ = "STUB: not implemented"; return }

func (m *SortedMap) Set(kv KV) error { _ = "STUB: not implemented"; return nil }

func (m *SortedMap) Get(key any) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

func (m *SortedMap) GetOr(key, dft any) any { _ = "STUB: not implemented"; return *new(any) }

func (m *SortedMap) GetString(key any) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (m *SortedMap) GetStringOr(key any, dft string) string { _ = "STUB: not implemented"; return "" }

func (m *SortedMap) HasKey(key any) bool { _ = "STUB: not implemented"; return false }

func (m *SortedMap) HasValue(value any) bool { _ = "STUB: not implemented"; return false }

func (m *SortedMap) Keys() []any { _ = "STUB: not implemented"; return nil }

func (m *SortedMap) Values() []any { _ = "STUB: not implemented"; return nil }

func (m *SortedMap) Range(iterator func(key, value any)) { _ = "STUB: not implemented"; return }

func (m *SortedMap) RangeIf(iterator func(key, value any) bool) { _ = "STUB: not implemented"; return }

func (m *SortedMap) Remove(key any) (value any, ok bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (m *SortedMap) Insert(sm *SortedMap) { _ = "STUB: not implemented"; return }

func (m *SortedMap) Copy() *SortedMap { _ = "STUB: not implemented"; return nil }

func (m *SortedMap) Format() []string { _ = "STUB: not implemented"; return nil }

func (m *SortedMap) Reset() { _ = "STUB: not implemented"; return }
