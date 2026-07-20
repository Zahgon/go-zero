package syncx

import "sync"

type (
	SingleFlight interface {
		Do(key string, fn func() (any, error)) (any, error)
		DoEx(key string, fn func() (any, error)) (any, bool, error)
	}

	call struct {
		wg  sync.WaitGroup
		val any
		err error
	}

	flightGroup struct {
		calls map[string]*call
		lock  sync.Mutex
	}
)

func NewSingleFlight() SingleFlight { _ = "STUB: not implemented"; return *new(SingleFlight) }

func (g *flightGroup) Do(key string, fn func() (any, error)) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (g *flightGroup) DoEx(key string, fn func() (any, error)) (val any, fresh bool, err error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

func (g *flightGroup) createCall(key string) (c *call, done bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (g *flightGroup) makeCall(c *call, key string, fn func() (any, error)) {
	_ = "STUB: not implemented"
	return
}
