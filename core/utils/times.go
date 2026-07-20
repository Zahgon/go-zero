package utils

import (
	"time"
)

type ElapsedTimer struct {
	start time.Duration
}

func NewElapsedTimer() *ElapsedTimer { _ = "STUB: not implemented"; return nil }

func (et *ElapsedTimer) Duration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (et *ElapsedTimer) Elapsed() string { _ = "STUB: not implemented"; return "" }

func (et *ElapsedTimer) ElapsedMs() string { _ = "STUB: not implemented"; return "" }

func CurrentMicros() int64 { _ = "STUB: not implemented"; return 0 }

func CurrentMillis() int64 { _ = "STUB: not implemented"; return 0 }
