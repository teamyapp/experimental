package routing

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/oauth"
	"github.com/teamyapp/experimental/yibolu/identity/app/pubsub"
	"net/http"

	"github.com/teamyapp/experimental/yibolu/identity/app/dao"
	"github.com/teamyapp/experimental/yibolu/identity/app/idgen"
	"github.com/teamyapp/experimental/yibolu/identity/app/security"
	"github.com/teamyapp/experimental/yibolu/identity/app/service"
)

type route struct {
	path       string
	method     string
	handleFunc http.HandlerFunc
}

func getRoutes(
	oauthProviders []oauth.OAuth,
	idGenerator idgen.IDGenerator,
	userDao dao.User,
	externalUserDao dao.ExternalUser,
	jwtAuthority security.JWTAuthority,
	caesarCipher security.CaesarCipher,
	pubsub pubsub.PubSub) []route {

	identity := service.NewIdentity(oauthProviders, idGenerator, userDao, externalUserDao, jwtAuthority, caesarCipher, pubsub)
	return []route{
		{
			path:       "/sign-in/{oauth_provider}",
			method:     http.MethodGet,
			handleFunc: newSignInHandlerFunc(identity),
		},
		{
			path:       "/sign-in/{oauth_provider}/callback/clients/{encrypted_client_id}",
			method:     http.MethodGet,
			handleFunc: newSignInFinishHandlerFunc(identity),
		},
		{
			path:       "/sign-in/clientID",
			method:     http.MethodGet,
			handleFunc: newGetClientIDHandlerFunc(identity),
		},
		{
			path:       "/sign-in/subscribe",
			method:     http.MethodGet,
			handleFunc: newSubscribeClientHandlerFunc(identity),
		},
	}
}
