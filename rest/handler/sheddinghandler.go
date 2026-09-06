package handler

import (
	"net/http"
	"sync"

	"github.com/zeromicro/go-zero/core/load"
	"github.com/zeromicro/go-zero/core/stat"
)

const serviceType = "api"

var (
	sheddingStat *load.SheddingStat
	lock         sync.Mutex
)

func SheddingHandler(shedder load.Shedder, metrics *stat.Metrics) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func ensureSheddingStat() { _ = "STUB: not implemented"; return }
