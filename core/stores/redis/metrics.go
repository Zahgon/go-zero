package redis

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	red "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/metric"
)

const namespace = "redis_client"

var (
	metricReqDur = metric.NewHistogramVec(&metric.HistogramVecOpts{
		Namespace: namespace,
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "redis client requests duration(ms).",
		Labels:    []string{"command"},
		Buckets:   []float64{0.25, 0.5, 1, 1.5, 2, 3, 5, 10, 25, 50, 100, 250, 500, 1000, 2000, 5000, 10000, 15000},
	})
	metricReqErr = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: namespace,
		Subsystem: "requests",
		Name:      "error_total",
		Help:      "redis client requests error count.",
		Labels:    []string{"command", "error"},
	})
	metricSlowCount = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: namespace,
		Subsystem: "requests",
		Name:      "slow_total",
		Help:      "redis client requests slow count.",
		Labels:    []string{"command"},
	})

	connLabels                         = []string{"key", "client_type"}
	connCollector                      = newCollector()
	_             prometheus.Collector = (*collector)(nil)
)

type (
	statGetter struct {
		clientType string
		key        string
		poolSize   int
		poolStats  func() *red.PoolStats
	}

	collector struct {
		hitDesc     *prometheus.Desc
		missDesc    *prometheus.Desc
		timeoutDesc *prometheus.Desc
		totalDesc   *prometheus.Desc
		idleDesc    *prometheus.Desc
		staleDesc   *prometheus.Desc
		maxDesc     *prometheus.Desc

		clients []*statGetter
		lock    sync.Mutex
	}
)

func newCollector() *collector { _ = "STUB: not implemented"; return nil }

func (s *collector) Describe(descs chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }

func (s *collector) Collect(metrics chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

func (s *collector) registerClient(client *statGetter) { _ = "STUB: not implemented"; return }
