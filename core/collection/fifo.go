package collection

import "sync"

const queueGrowThreshold = 256

type Queue struct {
	lock     sync.Mutex
	elements []any
	head     int
	tail     int
	count    int
}

func NewQueue(size int) *Queue { _ = "STUB: not implemented"; return nil }

func (q *Queue) Empty() bool { _ = "STUB: not implemented"; return false }

func (q *Queue) Put(element any) { _ = "STUB: not implemented"; return }

func (q *Queue) Take() (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

func nextQueueCapacity(capacity int) int { _ = "STUB: not implemented"; return 0 }
