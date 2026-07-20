package hash

import (
	"sync"

	"github.com/zeromicro/go-zero/core/lang"
)

const (
	TopWeight = 100

	minReplicas = 100
	prime       = 16777619
)

type (
	Func func(data []byte) uint64

	ConsistentHash struct {
		hashFunc Func
		replicas int
		keys     []uint64
		ring     map[uint64][]any
		nodes    map[string]lang.PlaceholderType
		lock     sync.RWMutex
	}
)

func NewConsistentHash() *ConsistentHash { _ = "STUB: not implemented"; return nil }

func NewCustomConsistentHash(replicas int, fn Func) *ConsistentHash {
	_ = "STUB: not implemented"
	return nil
}

func (h *ConsistentHash) Add(node any) { _ = "STUB: not implemented"; return }

func (h *ConsistentHash) AddWithReplicas(node any, replicas int) { _ = "STUB: not implemented"; return }

func (h *ConsistentHash) AddWithWeight(node any, weight int) { _ = "STUB: not implemented"; return }

func (h *ConsistentHash) Get(v any) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

func (h *ConsistentHash) Remove(node any) { _ = "STUB: not implemented"; return }

func (h *ConsistentHash) removeRingNode(hash uint64, nodeRepr string) {
	_ = "STUB: not implemented"
	return
}

func (h *ConsistentHash) addNode(nodeRepr string) { _ = "STUB: not implemented"; return }

func (h *ConsistentHash) containsNode(nodeRepr string) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *ConsistentHash) removeNode(nodeRepr string) { _ = "STUB: not implemented"; return }

func innerRepr(node any) string { _ = "STUB: not implemented"; return "" }

func repr(node any) string { _ = "STUB: not implemented"; return "" }
