package stringx

import (
	"math/rand"
	"sync"
	"time"
)

const (
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits  = 6
	idLen          = 8
	defaultRandLen = 8
	letterIdxMask  = 1<<letterIdxBits - 1
	letterIdxMax   = 63 / letterIdxBits
)

var src = newLockedSource(time.Now().UnixNano())

type lockedSource struct {
	source rand.Source
	lock   sync.Mutex
}

func newLockedSource(seed int64) *lockedSource { _ = "STUB: not implemented"; return nil }

func (ls *lockedSource) Int63() int64 { _ = "STUB: not implemented"; return 0 }

func (ls *lockedSource) Seed(seed int64) { _ = "STUB: not implemented"; return }

func Rand() string { _ = "STUB: not implemented"; return "" }

func RandId() string { _ = "STUB: not implemented"; return "" }

func Randn(n int) string { _ = "STUB: not implemented"; return "" }

func Seed(seed int64) { _ = "STUB: not implemented"; return }
