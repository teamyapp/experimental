package pubsub

type PubSub interface {
	Subscribe(topic string, callback func(data interface{})) Subscription
	Publish(topic string, data interface{}) error
	Start()
	Stop()
}

type Subscription interface {
	Unsubscribe() error
}

type RabbitMQ struct {

}
