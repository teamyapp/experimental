package queue

type MessageQueue interface {
	Subscribe(topic string, callback func(message string)) error
	Publish(topic, message string)
}

