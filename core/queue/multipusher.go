package queue

type MultiPusher struct {
	name    string
	pushers []Pusher
}

func NewMultiPusher(pushers []Pusher) Pusher { _ = "STUB: not implemented"; return *new(Pusher) }

func (pusher *MultiPusher) Name() string { _ = "STUB: not implemented"; return "" }

func (pusher *MultiPusher) Push(message string) error { _ = "STUB: not implemented"; return nil }
