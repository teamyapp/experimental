package routing

import (
	"github.com/google/uuid"
	"github.com/teamyapp/experimental/yibolu/identity/app/entity"
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
		state := entity.OAuthState{
			StateID: uuid.New().String(),
			ClientID: clientId,
			OAuthProvider: oauthProviderName,
		}
		identity.StateManager.SaveOAuthState(state)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func newSignInFinishHandlerFunc(identity service.Identity) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stateID := r.FormValue("state")
		state := identity.StateManager.GetOAuthCallbackState(stateID)
		authCode := r.FormValue("code")
		identity.FinishOAuthSignIn(authCode, state)
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