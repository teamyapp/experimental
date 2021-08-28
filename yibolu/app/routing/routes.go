package routing

import (
	"net/http"

	"github.com/teamyapp/identity/app/dao"
	"github.com/teamyapp/identity/app/idgen"
	"github.com/teamyapp/identity/app/security"
	"github.com/teamyapp/identity/app/service"
)

type route struct {
	path       string
	method     string
	handleFunc http.HandlerFunc
}

func getRoutes(
	idGenerator idgen.IDGenerator,
	userDao dao.User,
	externalUserDao dao.ExternalUser,
	jwtAuthority security.JWTAuthority,
	caesarCipher security.CaesarCipher) []route {

	authenticationService := service.NewIdentity(idGenerator, userDao, externalUserDao, jwtAuthority, caesarCipher)
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
	}
}
