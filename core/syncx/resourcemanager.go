package syncx

import (
	"io"
	"sync"
)

type ResourceManager struct {
	resources    map[string]io.Closer
	singleFlight SingleFlight
	lock         sync.RWMutex
}

func NewResourceManager() *ResourceManager { _ = "STUB: not implemented"; return nil }

func (manager *ResourceManager) Close() error { _ = "STUB: not implemented"; return nil }

func (manager *ResourceManager) GetResource(key string, create func() (io.Closer, error)) (
	io.Closer, error) {
	_ = "STUB: not implemented"
	return *new(io.Closer), nil
}

func (manager *ResourceManager) Inject(key string, resource io.Closer) {
	_ = "STUB: not implemented"
	return
}
