package routing

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/service"
	"net/http"
)

var clientIdKey = "clientId"

func newSignInHandlerFunc(identity service.Identity) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		oauthProviderName := r.FormValue("oauth_provider")
		clientId := r.FormValue(clientIdKey)

		url, err := identity.RequestOAuthSignInURL(oauthProviderName, clientId)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func newSignInFinishHandlerFunc(identity service.Identity) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: research gorilla router to get the path parameter
		providerName := r.FormValue("oauth_provider")
		oauth, err := identity.GetOAuthProvider(providerName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		stateID, err := oauth.GetStateID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		authCode, err := oauth.GetAuthorizationCode(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		identity.FinishOAuthSignIn(authCode, stateID, providerName)
	}
}

func newGetClientIDHandlerFunc(identity service.Identity) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("not implemented")
	}
}

func newSubscribeClientHandlerFunc(identity service.Identity) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("not implemented")
	}
}