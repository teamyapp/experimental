package oauth

import "net/http"

type OAuth interface {
	GetName() string
	RedirectToLogin(w http.ResponseWriter, r *http.Request)
	GetUserInfo(w http.ResponseWriter, r *http.Request)
}
