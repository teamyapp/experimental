package oauth

type OAuth interface {
	GetName() string
	HandleLogin()
	HandleCallback()
}
