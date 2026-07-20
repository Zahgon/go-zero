package migrate

import (
	"net/http"
	"time"
)

var client = http.Client{
	Timeout: 5 * time.Second,
}

func getLatest(repo string, verbose bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
