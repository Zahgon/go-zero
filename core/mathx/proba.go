package mathx

import (
	"math/rand"
	"sync"
)

type Proba struct {
	r    *rand.Rand
	lock sync.Mutex
}

func NewProba() *Proba { _ = "STUB: not implemented"; return nil }

func (p *Proba) TrueOnProba(proba float64) (truth bool) { _ = "STUB: not implemented"; return false }
