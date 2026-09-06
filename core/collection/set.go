package collection

import "github.com/zeromicro/go-zero/core/lang"

type Set[T comparable] struct {
	data map[T]lang.PlaceholderType
}

func NewSet[T comparable]() *Set[T] { _ = "STUB: not implemented"; return nil }

func (s *Set[T]) Add(items ...T) { _ = "STUB: not implemented"; return }

func (s *Set[T]) Clear() { _ = "STUB: not implemented"; return }

func (s *Set[T]) Contains(item T) bool { _ = "STUB: not implemented"; return false }

func (s *Set[T]) Count() int { _ = "STUB: not implemented"; return 0 }

func (s *Set[T]) Keys() []T { _ = "STUB: not implemented"; return nil }

func (s *Set[T]) Remove(item T) { _ = "STUB: not implemented"; return }
