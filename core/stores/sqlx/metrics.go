package sqlx

import (
	"database/sql"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/zeromicro/go-zero/core/metric"
)

const namespace = "sql_client"

var (
	metricReqDur = metric.NewHistogramVec(&metric.HistogramVecOpts{
		Namespace: namespace,
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "mysql client requests duration(ms).",
		Labels:    []string{"command"},
		Buckets:   []float64{0.25, 0.5, 1, 1.5, 2, 3, 5, 10, 25, 50, 100, 250, 500, 1000, 2000, 5000, 10000, 15000},
	})
	metricReqErr = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: namespace,
		Subsystem: "requests",
		Name:      "error_total",
		Help:      "mysql client requests error count.",
		Labels:    []string{"command", "error"},
	})
	metricSlowCount = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: namespace,
		Subsystem: "requests",
		Name:      "slow_total",
		Help:      "mysql client requests slow count.",
		Labels:    []string{"command"},
	})

	connLabels                         = []string{"db_name", "hash"}
	connCollector                      = newCollector()
	_             prometheus.Collector = (*collector)(nil)
)

type (
	statGetter struct {
		host      string
		dbName    string
		hash      string
		poolStats func() sql.DBStats
	}

	collector struct {
		maxOpenConnections *prometheus.Desc

		openConnections  *prometheus.Desc
		inUseConnections *prometheus.Desc
		idleConnections  *prometheus.Desc

		waitCount         *prometheus.Desc
		waitDuration      *prometheus.Desc
		maxIdleClosed     *prometheus.Desc
		maxIdleTimeClosed *prometheus.Desc
		maxLifetimeClosed *prometheus.Desc

		clients []*statGetter
		lock    sync.Mutex
	}
)

func newCollector() *collector { _ = "STUB: not implemented"; return nil }

func (c *collector) Describe(ch chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }

func (c *collector) Collect(ch chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

func (c *collector) registerClient(client *statGetter) { _ = "STUB: not implemented"; return }
