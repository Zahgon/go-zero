package queue

type (
	Consumer interface {
		Consume(string) error
		OnEvent(event any)
	}

	ConsumerFactory func() (Consumer, error)
)
