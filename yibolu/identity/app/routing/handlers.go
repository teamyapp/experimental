package routing

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/oauth"
	"net/http"

	"github.com/teamyapp/experimental/yibolu/identity/app/service"
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
		state := oauth.GetOAuthCallbackState(r.FormValue("state"))
		authCode := r.FormValue("code")
		identity.FinishOAuthSignIn(authCode, state.OauthProvider, state.ClientId)
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