package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/stat"
)

func MetricHandler(metrics *stat.Metrics) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
