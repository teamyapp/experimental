package queue

import (
	"errors"
	"time"
)

type ChannelQueue struct {
	// key: client id, value: channel for jwt
	queue			map[string] chan string
}

func (q ChannelQueue) Subscribe(clientID string, callback func()) error {
	if _, ok := q.queue[clientID]; ok {
		return errors.New("already subscribed client: " + clientID)
	}
	q.queue[clientID] = make(chan string)
	go q.waitForMessage(topic, callback)
	return nil
}

func (q ChannelQueue) waitForMessage(topic string, callback func(message string))  {
	select {
	case message := <-q.queue[topic]:
		callback(message)
		return nil
	case <- time.After(time.Minute * 5):
		delete(q.queue, clientID)
		return errors.New("timeout when getting JWT")
	}
}


func (q ChannelQueue) SendJWT(clientId, jwt string) error {
	if val, ok := q.queue[clientId]; ok {
		val <- jwt

		return nil
	}
	return errors.New("cannot find clientId")
}

func NewChannelQueue() ChannelQueue {
	return ChannelQueue{
		queue:		make(map[string] chan string),
	}
}
