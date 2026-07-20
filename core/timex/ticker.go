package timex

import (
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/lang"
)

var errTimeout = errors.New("timeout")

type (
	Ticker interface {
		Chan() <-chan time.Time
		Stop()
	}

	FakeTicker interface {
		Ticker
		Done()
		Tick()
		Wait(d time.Duration) error
	}

	fakeTicker struct {
		c    chan time.Time
		done chan lang.PlaceholderType
	}

	realTicker struct {
		*time.Ticker
	}
)

func NewTicker(d time.Duration) Ticker { _ = "STUB: not implemented"; return *new(Ticker) }

func (rt *realTicker) Chan() <-chan time.Time { _ = "STUB: not implemented"; return nil }

func NewFakeTicker() FakeTicker { _ = "STUB: not implemented"; return *new(FakeTicker) }

func (ft *fakeTicker) Chan() <-chan time.Time { _ = "STUB: not implemented"; return nil }

func (ft *fakeTicker) Done() { _ = "STUB: not implemented"; return }

func (ft *fakeTicker) Stop() { _ = "STUB: not implemented"; return }

func (ft *fakeTicker) Tick() { _ = "STUB: not implemented"; return }

func (ft *fakeTicker) Wait(d time.Duration) error { _ = "STUB: not implemented"; return nil }
