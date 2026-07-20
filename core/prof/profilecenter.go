package prof

import (
	"sync"
	"time"
)

type (
	profileSlot struct {
		lifecount int64
		lastcount int64
		lifecycle int64
		lastcycle int64
	}

	profileCenter struct {
		lock  sync.RWMutex
		slots map[string]*profileSlot
	}
)

const flushInterval = 5 * time.Minute

var pc = &profileCenter{
	slots: make(map[string]*profileSlot),
}

func init() {
	flushRepeatedly()
}

func flushRepeatedly() { _ = "STUB: not implemented"; return }

func report(name string, duration time.Duration) { _ = "STUB: not implemented"; return }

func loadOrStoreSlot(name string, duration time.Duration) *profileSlot {
	_ = "STUB: not implemented"
	return nil
}

func generateReport() string { _ = "STUB: not implemented"; return "" }
