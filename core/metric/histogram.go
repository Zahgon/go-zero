package metric

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

type (
	HistogramVecOpts struct {
		Namespace   string
		Subsystem   string
		Name        string
		Help        string
		Labels      []string
		Buckets     []float64
		ConstLabels map[string]string
	}

	HistogramVec interface {
		Observe(v int64, labels ...string)

		ObserveFloat(v float64, labels ...string)
		close() bool
	}

	promHistogramVec struct {
		histogram *prom.HistogramVec
	}
)

func NewHistogramVec(cfg *HistogramVecOpts) HistogramVec {
	_ = "STUB: not implemented"
	return *new(HistogramVec)
}

func (hv *promHistogramVec) Observe(v int64, labels ...string) { _ = "STUB: not implemented"; return }

func (hv *promHistogramVec) ObserveFloat(v float64, labels ...string) {
	_ = "STUB: not implemented"
	return
}

func (hv *promHistogramVec) close() bool { _ = "STUB: not implemented"; return false }
