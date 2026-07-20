package load

import (
	"time"
)

type (
	SheddingStat struct {
		name  string
		total int64
		pass  int64
		drop  int64
	}

	snapshot struct {
		Total int64
		Pass  int64
		Drop  int64
	}
)

func NewSheddingStat(name string) *SheddingStat { _ = "STUB: not implemented"; return nil }

func (s *SheddingStat) IncrementTotal() { _ = "STUB: not implemented"; return }

func (s *SheddingStat) IncrementPass() { _ = "STUB: not implemented"; return }

func (s *SheddingStat) IncrementDrop() { _ = "STUB: not implemented"; return }

func (s *SheddingStat) loop(c <-chan time.Time) { _ = "STUB: not implemented"; return }

func (s *SheddingStat) reset() snapshot { _ = "STUB: not implemented"; return *new(snapshot) }

func (s *SheddingStat) run() { _ = "STUB: not implemented"; return }
