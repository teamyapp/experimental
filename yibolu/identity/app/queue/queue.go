package queue

type MessageQueue interface {
	Subscribe(clientID string) error
	GetJWT(clientID string, callback func(jwt string)) error
}

