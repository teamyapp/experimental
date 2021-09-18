package queue

import (
	"errors"
	"time"
)

type ChannelQueue struct {
	queue			map[string] chan string
}

func (q ChannelQueue) Subscribe(clientID string) error {
	if _, ok := q.queue[clientID]; ok {
		return errors.New("already subscribed client: " + clientID)
	}
	q.queue[clientID] = make(chan string)
	return nil
}

func (q ChannelQueue) GetJWT(clientID string, callback func(jwt string)) error {
	start := time.Now()
	for {
		select {
		case jwt := <-q.queue[clientID]:
			callback(jwt)
			return nil

		default:
			now := time.Now()
			diff := now.Sub(start)
			if diff.Seconds() > 300 { //5 mins
				delete(q.queue, clientID)
				return errors.New("timeout when getting JWT")
			}
		}
	}
}

func NewChannelQueue() ChannelQueue {
	return ChannelQueue{
		queue:		make(map[string] chan string),
	}
}
