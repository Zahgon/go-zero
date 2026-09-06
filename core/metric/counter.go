package metric

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

type (
	CounterVecOpts VectorOpts

	CounterVec interface {
		Inc(labels ...string)

		Add(v float64, labels ...string)
		close() bool
	}

	promCounterVec struct {
		counter *prom.CounterVec
	}
)

func NewCounterVec(cfg *CounterVecOpts) CounterVec {
	_ = "STUB: not implemented"
	return *new(CounterVec)
}

func (cv *promCounterVec) Add(v float64, labels ...string) { _ = "STUB: not implemented"; return }

func (cv *promCounterVec) Inc(labels ...string) { _ = "STUB: not implemented"; return }

func (cv *promCounterVec) close() bool { _ = "STUB: not implemented"; return false }
