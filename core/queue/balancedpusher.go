package queue

import (
	"errors"
)

var ErrNoAvailablePusher = errors.New("no available pusher")

type BalancedPusher struct {
	name    string
	pushers []Pusher
	index   uint64
}

func NewBalancedPusher(pushers []Pusher) Pusher { _ = "STUB: not implemented"; return *new(Pusher) }

func (pusher *BalancedPusher) Name() string { _ = "STUB: not implemented"; return "" }

func (pusher *BalancedPusher) Push(message string) error { _ = "STUB: not implemented"; return nil }
