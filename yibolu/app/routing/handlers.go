package routing

import (
	"net/http"

	"github.com/teamyapp/identity/app/service"
)

func newSignInHandlerFunc(authenticationService service.Authentication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("not implemented")
	}
}

func newSignInFinishHandlerFunc(authenticationService service.Authentication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("not implemented")
	}
}
