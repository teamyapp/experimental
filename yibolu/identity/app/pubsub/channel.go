package pubsub

import (
	"errors"
)

type ChannelPubSub struct {
	started       	bool
	sub				map[string]*ChannelQueue
}

func (c *ChannelPubSub) Start() {
	c.started = true
}

func (c *ChannelPubSub) Stop() {
	c.started = false
}

func (c *ChannelPubSub) Subscribe(topic string, callback func(data interface{})) Subscription {
	sub := ChannelQueue{
		pubSub:   c,
		topic:    topic,
		callback: callback,
	}
	c.sub[topic] = &sub
	return sub
}

func (c *ChannelPubSub) Publish(topic string, data interface{}) error {
	if !c.started {
		return errors.New("pubSub not started")
	}

	sub, ok := c.sub[topic]
	if !ok {
		return nil
	}

	sub.callback(data)
	return nil
}

func NewChannelPubSub() *ChannelPubSub {
	return &ChannelPubSub{
		sub: make(map[string]*ChannelQueue),
	}
}

var _ PubSub = (*ChannelPubSub)(nil)

type ChannelQueue struct {
	pubSub   *ChannelPubSub
	topic    string
	callback func(data interface{})
}

func (c ChannelQueue) Unsubscribe() error {
	delete(c.pubSub.sub, c.topic)
	return nil
}

var _ Subscription = (*ChannelQueue)(nil)