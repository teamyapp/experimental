package registry

type ClientRegistry interface {
	AssignClient(userID string, clientID string) error
	RemoveClient(userID string, clientID string) error
	RequestClientID() (string, error)
}
