package routing

import (
	"net/http"

	"github.com/teamyapp/experimental/yibolu/identity/app/service"
)

func newSignInHandlerFunc(authenticationService service.Authentication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		oauthProviderName := r.FormValue("oauth_provider")
		clientId := r.FormValue("clientId")

		authenticationService.RequestOAuthSignIn(oauthProviderName, clientId)
	}
}

func newSignInFinishHandlerFunc(authenticationService service.Authentication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("not implemented")
	}
}

func newGetClientIDHandlerFunc(authenticationService service.Authentication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("not implemented")
	}
}

func newSubscribeClientHandlerFunc(authenticationService service.Authentication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("not implemented")
	}
}