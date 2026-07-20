package mapping

type (
	Valuer interface {
		Value(key string) (any, bool)
	}

	valuerWithParent interface {
		Valuer

		Parent() valuerWithParent
	}

	node struct {
		current Valuer
		parent  valuerWithParent
	}

	valueWithParent struct {
		value  any
		parent valuerWithParent
	}

	mapValuer map[string]any

	simpleValuer node

	recursiveValuer node
)

func (mv mapValuer) Value(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (sv simpleValuer) Value(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (sv simpleValuer) Parent() valuerWithParent {
	_ = "STUB: not implemented"
	return *new(valuerWithParent)
}

func (rv recursiveValuer) Value(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (rv recursiveValuer) Parent() valuerWithParent {
	_ = "STUB: not implemented"
	return *new(valuerWithParent)
}
