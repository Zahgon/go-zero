package conf

import (
	"sync"
)

type PropertyError struct {
	message string
}

type Properties interface {
	GetString(key string) string
	SetString(key, value string)
	GetInt(key string) int
	SetInt(key string, value int)
	ToString() string
}

type mapBasedProperties struct {
	properties map[string]string
	lock       sync.RWMutex
}

func LoadProperties(filename string, opts ...Option) (Properties, error) {
	_ = "STUB: not implemented"
	return *new(Properties), nil
}

func (config *mapBasedProperties) GetString(key string) string {
	_ = "STUB: not implemented"
	return ""
}

func (config *mapBasedProperties) SetString(key, value string) { _ = "STUB: not implemented"; return }

func (config *mapBasedProperties) GetInt(key string) int { _ = "STUB: not implemented"; return 0 }

func (config *mapBasedProperties) SetInt(key string, value int) { _ = "STUB: not implemented"; return }

func (config *mapBasedProperties) ToString() string { _ = "STUB: not implemented"; return "" }

func (configError *PropertyError) Error() string { _ = "STUB: not implemented"; return "" }

func NewProperties() Properties { _ = "STUB: not implemented"; return *new(Properties) }
