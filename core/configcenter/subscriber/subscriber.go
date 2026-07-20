package subscriber

type Subscriber interface {
	AddListener(listener func()) error

	Value() (string, error)
}
