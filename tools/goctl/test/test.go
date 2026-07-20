package test

import (
	"testing"
)

type Data[T, Y any] struct {
	Name  string
	Input T
	Want  Y
	E     error
}

type Option[T, Y any] func(*Executor[T, Y])
type assertFn[Y any] func(t *testing.T, expected, actual Y) bool

func WithComparison[T, Y any](comparisonFn assertFn[Y]) Option[T, Y] {
	_ = "STUB: not implemented"
	return nil
}

type Executor[T, Y any] struct {
	list    []Data[T, Y]
	equalFn assertFn[Y]
}

func NewExecutor[T, Y any](opt ...Option[T, Y]) *Executor[T, Y] {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor[T, Y]) Add(data ...Data[T, Y]) { _ = "STUB: not implemented"; return }

func (e *Executor[T, Y]) Run(t *testing.T, do func(T) Y) { _ = "STUB: not implemented"; return }

func (e *Executor[T, Y]) RunE(t *testing.T, do func(T) (Y, error)) {
	_ = "STUB: not implemented"
	return
}
