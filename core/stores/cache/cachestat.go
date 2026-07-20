package cache

import (
	"time"

	"github.com/zeromicro/go-zero/core/timex"
)

const statInterval = time.Minute

type Stat struct {
	name string

	Total   uint64
	Hit     uint64
	Miss    uint64
	DbFails uint64
}

func NewStat(name string) *Stat { _ = "STUB: not implemented"; return nil }

func (s *Stat) IncrementTotal() { _ = "STUB: not implemented"; return }

func (s *Stat) IncrementHit() { _ = "STUB: not implemented"; return }

func (s *Stat) IncrementMiss() { _ = "STUB: not implemented"; return }

func (s *Stat) IncrementDbFails() { _ = "STUB: not implemented"; return }

func (s *Stat) statLoop(ticker timex.Ticker) { _ = "STUB: not implemented"; return }
