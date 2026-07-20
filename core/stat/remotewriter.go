package stat

import (
	"errors"
	"time"
)

const (
	httpTimeout     = time.Second * 5
	jsonContentType = "application/json; charset=utf-8"
)

var ErrWriteFailed = errors.New("submit failed")

type RemoteWriter struct {
	endpoint string
}

func NewRemoteWriter(endpoint string) Writer { _ = "STUB: not implemented"; return *new(Writer) }

func (rw *RemoteWriter) Write(report *StatReport) error { _ = "STUB: not implemented"; return nil }
