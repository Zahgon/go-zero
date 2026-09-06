package stat

import (
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/executors"
	"github.com/zeromicro/go-zero/core/syncx"
)

var (
	logInterval  = time.Minute
	writerLock   sync.Mutex
	reportWriter Writer = nil
	logEnabled          = syncx.ForAtomicBool(true)
)

type (
	Writer interface {
		Write(report *StatReport) error
	}

	StatReport struct {
		Name          string  `json:"name"`
		Timestamp     int64   `json:"tm"`
		Pid           int     `json:"pid"`
		ReqsPerSecond float32 `json:"qps"`
		Drops         int     `json:"drops"`
		Average       float32 `json:"avg"`
		Median        float32 `json:"med"`
		Top90th       float32 `json:"t90"`
		Top99th       float32 `json:"t99"`
		Top99p9th     float32 `json:"t99p9"`
	}

	Metrics struct {
		executor  *executors.PeriodicalExecutor
		container *metricsContainer
	}
)

func DisableLog() { _ = "STUB: not implemented"; return }

func SetReportWriter(writer Writer) { _ = "STUB: not implemented"; return }

func NewMetrics(name string) *Metrics { _ = "STUB: not implemented"; return nil }

func (m *Metrics) Add(task Task) { _ = "STUB: not implemented"; return }

func (m *Metrics) AddDrop() { _ = "STUB: not implemented"; return }

func (m *Metrics) SetName(name string) { _ = "STUB: not implemented"; return }

type (
	tasksDurationPair struct {
		tasks    []Task
		duration time.Duration
		drops    int
	}

	metricsContainer struct {
		name     string
		pid      int
		tasks    []Task
		duration time.Duration
		drops    int
	}
)

func (c *metricsContainer) AddTask(v any) bool { _ = "STUB: not implemented"; return false }

func (c *metricsContainer) Execute(v any) { _ = "STUB: not implemented"; return }

func (c *metricsContainer) RemoveAll() any { _ = "STUB: not implemented"; return *new(any) }

func getTopDuration(tasks []Task) float32 { _ = "STUB: not implemented"; return 0 }

func log(report *StatReport) { _ = "STUB: not implemented"; return }

func writeReport(report *StatReport) { _ = "STUB: not implemented"; return }
