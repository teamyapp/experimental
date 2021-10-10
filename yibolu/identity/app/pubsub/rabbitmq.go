package pubsub

import (
	"errors"
	"sync"
)

type RabbitMQ struct {
	started       	bool
	mutex         	sync.Mutex
	sub				map[string]*RabbitQueue
}

func (r *RabbitMQ) Start() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.started = true
}

func (r *RabbitMQ) Stop() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.started = false
}

func (r *RabbitMQ) Subscribe(topic string, callback func(data interface{})) Subscription {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	sub := RabbitQueue{
		pubSub:   r,
		topic:    topic,
		callback: callback,
	}
	r.sub[topic] = &sub
	return sub
}

func (r *RabbitMQ) Publish(topic string, data interface{}) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if !r.started {
		return errors.New("pubSub not started")
	}

	sub, ok := r.sub[topic]
	if !ok {
		return nil
	}

	sub.callback(data)
	return nil
}

func NewRabbitMQ() *RabbitMQ {
	return &RabbitMQ{
		mutex:         	sync.Mutex{},
		sub: 			make(map[string]*RabbitQueue),
	}
}

var _ PubSub = (*RabbitMQ)(nil)

type RabbitQueue struct {
	pubSub   *RabbitMQ
	topic    string
	callback func(data interface{})
}

func (r RabbitQueue) Unsubscribe() error {
	delete(r.pubSub.sub, r.topic)
	return nil
}

var _ Subscription = (*RabbitQueue)(nil)