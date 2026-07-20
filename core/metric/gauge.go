package metric

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

type (
	GaugeVecOpts VectorOpts

	GaugeVec interface {
		Set(v float64, labels ...string)

		Inc(labels ...string)

		Dec(labels ...string)

		Add(v float64, labels ...string)

		Sub(v float64, labels ...string)
		close() bool
	}

	promGaugeVec struct {
		gauge *prom.GaugeVec
	}
)

func NewGaugeVec(cfg *GaugeVecOpts) GaugeVec { _ = "STUB: not implemented"; return *new(GaugeVec) }

func (gv *promGaugeVec) Add(v float64, labels ...string) { _ = "STUB: not implemented"; return }

func (gv *promGaugeVec) Dec(labels ...string) { _ = "STUB: not implemented"; return }

func (gv *promGaugeVec) Inc(labels ...string) { _ = "STUB: not implemented"; return }

func (gv *promGaugeVec) Set(v float64, labels ...string) { _ = "STUB: not implemented"; return }

func (gv *promGaugeVec) Sub(v float64, labels ...string) { _ = "STUB: not implemented"; return }

func (gv *promGaugeVec) close() bool { _ = "STUB: not implemented"; return false }
