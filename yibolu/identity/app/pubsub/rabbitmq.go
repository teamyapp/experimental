package pubsub

import (
	"errors"
	"sync"
)

type RabbitMQ struct {
	started      	bool
	mutex        	sync.Mutex
	subscriptions 	map[string][]*RabbitQueue
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
	subscription := &RabbitQueue{
		pubSub:   r,
		topic:    topic,
		callback: callback,
	}
	r.subscriptions[topic] = append(r.subscriptions[topic], subscription)
	return subscription
}

func (r *RabbitMQ) Publish(topic string, data interface{}) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if !r.started {
		return errors.New("pubSub not started")
	}

	subs, ok := r.subscriptions[topic]
	if !ok {
		return errors.New("we never subscribed the topic")
	}

	for _, sub := range subs {
		go func(sub *RabbitQueue) {
			sub.callback(data)
		}(sub)
	}

	return nil
}

func NewRabbitMQ() *RabbitMQ {
	return &RabbitMQ{
		subscriptions: make(map[string][]*RabbitQueue),
	}
}

var _ PubSub = (*RabbitMQ)(nil)

type RabbitQueue struct {
	pubSub   *RabbitMQ
	topic    string
	callback func(data interface{})
}

func (r *RabbitQueue) Unsubscribe() error {
	r.pubSub.mutex.Lock()
	defer r.pubSub.mutex.Unlock()
	subs := r.pubSub.subscriptions[r.topic]
	newSubs := make([]*RabbitQueue, 0)
	for _, sub := range subs {
		if sub == r {
			continue
		}
		newSubs = append(newSubs, sub)
	}
	if len(newSubs) == 0 {
		delete(r.pubSub.subscriptions, r.topic)
	} else {
		r.pubSub.subscriptions[r.topic] = newSubs
	}
	return nil
}

var _ Subscription = (*RabbitQueue)(nil)