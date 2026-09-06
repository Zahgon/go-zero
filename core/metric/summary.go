package metric

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

type (
	SummaryVecOpts struct {
		VecOpt     VectorOpts
		Objectives map[float64]float64
	}

	SummaryVec interface {
		Observe(v float64, labels ...string)
		close() bool
	}

	promSummaryVec struct {
		summary *prom.SummaryVec
	}
)

func NewSummaryVec(cfg *SummaryVecOpts) SummaryVec {
	_ = "STUB: not implemented"
	return *new(SummaryVec)
}

func (sv *promSummaryVec) Observe(v float64, labels ...string) { _ = "STUB: not implemented"; return }

func (sv *promSummaryVec) close() bool { _ = "STUB: not implemented"; return false }
