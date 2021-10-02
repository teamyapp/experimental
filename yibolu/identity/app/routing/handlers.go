package routing

import (
	"net/http"

	"github.com/teamyapp/experimental/yibolu/identity/app/service"
)

var clientIdKey = "clientId"

func newSignInHandlerFunc(authenticationService service.Authentication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		oauthProviderName := r.FormValue("oauth_provider")
		clientId := r.FormValue(clientIdKey)

		authenticationService.RequestOAuthSignIn(w, r, oauthProviderName, clientId)
	}
}

func newSignInFinishHandlerFunc(authenticationService service.Authentication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		oauthProvider := "google"
		clientId, _ := r.Cookie(clientIdKey)
		authenticationService.FinishOAuthSignIn(w, r, oauthProvider, clientId)
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