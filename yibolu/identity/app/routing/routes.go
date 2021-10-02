package routing

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/oauth"
	"github.com/teamyapp/experimental/yibolu/identity/app/queue"
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
	queue queue.MessageQueue) []route {

	authenticationService := service.NewIdentity(oauthProviders, idGenerator, userDao, externalUserDao, jwtAuthority, caesarCipher)
	return []route{
		{
			path:       "/sign-in/{oauth_provider}",
			method:     http.MethodGet,
			handleFunc: newSignInHandlerFunc(authenticationService),
		},
		{
			path:       "/sign-in/{oauth_provider}/callback/clients/{encrypted_client_id}",
			method:     http.MethodGet,
			handleFunc: newSignInFinishHandlerFunc(authenticationService),
		},
		{
			path:       "/sign-in/clientID",
			method:     http.MethodGet,
			handleFunc: newGetClientIDHandlerFunc(authenticationService),
		},
		{
			path:       "/sign-in/subscribe",
			method:     http.MethodGet,
			handleFunc: newSubscribeClientHandlerFunc(authenticationService),
		},
	}
}
