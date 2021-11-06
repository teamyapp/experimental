package routing

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/oauth"
	"github.com/teamyapp/experimental/yibolu/identity/app/pubsub"
	"net/http"

	"github.com/teamyapp/experimental/yibolu/identity/app/dao"
	"github.com/teamyapp/experimental/yibolu/identity/app/idgen"
	"github.com/teamyapp/experimental/yibolu/identity/app/security"

	"github.com/gorilla/mux"
)

func NewServer(
	oauthProviders []oauth.OAuth,
	idGenerator idgen.IDGenerator,
	userDao dao.User,
	externalUserDao dao.ExternalUser,
	jwtAuthority security.JWTAuthority,
	caesarCipher security.CaesarCipher,
	pubsub pubsub.PubSub,
	stateManager oauth.StateManager) *http.ServeMux {

	serveMux := http.NewServeMux()
	router := mux.NewRouter()
	routes := getRoutes(oauthProviders, idGenerator, userDao, externalUserDao, jwtAuthority, caesarCipher, pubsub, stateManager)
	for _, r := range routes {
		router.HandleFunc(r.path, r.handleFunc).Methods(r.method)
	}

	serveMux.HandleFunc("/", enableCORS(router.ServeHTTP))
	return serveMux
}

func enableCORS(handlerFunc http.HandlerFunc) http.HandlerFunc { // Closure
	return func(writer http.ResponseWriter, request *http.Request) { // Closure
		writer.Header().Set("Access-Control-Allow-Origin", "*")                                // Decorator
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE") // Decorator
		writer.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, Authorization") // Decorator
		if request.Method == http.MethodOptions { // Decorator
			return // Decorator
		}

		handlerFunc(writer, request) // Closure, Decorator
	}
}
