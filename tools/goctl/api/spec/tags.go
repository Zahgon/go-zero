package spec

import (
	"errors"
)

var errTagNotExist = errors.New("tag does not exist")

type (
	Tag struct {
		Key string

		Name string

		Options []string
	}

	Tags struct {
		tags []*Tag
	}
)

func Parse(tag string) (*Tags, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tags) Get(key string) (*Tag, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tags) Keys() []string { _ = "STUB: not implemented"; return nil }

func (t *Tags) Tags() []*Tag { _ = "STUB: not implemented"; return nil }
