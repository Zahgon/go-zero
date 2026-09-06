package queue

import (
	"sync"

	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/core/threading"
)

const queueName = "queue"

type (
	Queue struct {
		name                 string
		metrics              *stat.Metrics
		producerFactory      ProducerFactory
		producerRoutineGroup *threading.RoutineGroup
		consumerFactory      ConsumerFactory
		consumerRoutineGroup *threading.RoutineGroup
		producerCount        int
		consumerCount        int
		active               int32
		channel              chan string
		quit                 chan struct{}
		listeners            []Listener
		eventLock            sync.Mutex
		eventChannels        []chan any
	}

	Listener interface {
		OnPause()
		OnResume()
	}

	Poller interface {
		Name() string
		Poll() string
	}

	Pusher interface {
		Name() string
		Push(string) error
	}
)

func NewQueue(producerFactory ProducerFactory, consumerFactory ConsumerFactory) *Queue {
	_ = "STUB: not implemented"
	return nil
}

func (q *Queue) AddListener(listener Listener) { _ = "STUB: not implemented"; return }

func (q *Queue) Broadcast(message any) { _ = "STUB: not implemented"; return }

func (q *Queue) SetName(name string) { _ = "STUB: not implemented"; return }

func (q *Queue) SetNumConsumer(count int) { _ = "STUB: not implemented"; return }

func (q *Queue) SetNumProducer(count int) { _ = "STUB: not implemented"; return }

func (q *Queue) Start() { _ = "STUB: not implemented"; return }

func (q *Queue) Stop() { _ = "STUB: not implemented"; return }

func (q *Queue) consume(eventChan chan any) { _ = "STUB: not implemented"; return }

func (q *Queue) consumeOne(consumer Consumer, message string) { _ = "STUB: not implemented"; return }

func (q *Queue) pause() { _ = "STUB: not implemented"; return }

func (q *Queue) produce() { _ = "STUB: not implemented"; return }

func (q *Queue) produceOne(producer Producer) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (q *Queue) resume() { _ = "STUB: not implemented"; return }

func (q *Queue) startConsumers(number int) { _ = "STUB: not implemented"; return }

func (q *Queue) startProducers(number int) { _ = "STUB: not implemented"; return }

type routineListener struct {
	queue *Queue
}

func (rl routineListener) OnProducerPause() { _ = "STUB: not implemented"; return }

func (rl routineListener) OnProducerResume() { _ = "STUB: not implemented"; return }
