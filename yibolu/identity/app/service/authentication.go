package service

import (
	"fmt"
	"github.com/teamyapp/experimental/yibolu/identity/app/channel"
	"github.com/teamyapp/experimental/yibolu/identity/app/dao"
	"github.com/teamyapp/experimental/yibolu/identity/app/idgen"
	"github.com/teamyapp/experimental/yibolu/identity/app/oauth"
	"github.com/teamyapp/experimental/yibolu/identity/app/pubsub"
	"github.com/teamyapp/experimental/yibolu/identity/app/security"
	"log"
	"net/http"
	"time"
)

var timeWait = time.Second * 300 // 5 minutes

type Authentication interface {
	RequestOAuthSignIn(w http.ResponseWriter, r *http.Request, oauthProvider string, clientId string) error
	FinishOAuthSignIn(w http.ResponseWriter, r *http.Request, oauthProvider string, clientId string) error
	ValidateAuthToken(authToken string) error
}

type Identity struct {
	idGenerator     idgen.IDGenerator
	userDao         dao.User
	externalUserDao dao.ExternalUser
	jwtAuthority    security.JWTAuthority
	caesarCipher    security.CaesarCipher
	oauth           map[string] oauth.OAuth
	pubsub    		pubsub.PubSub
}

func (s Identity) RequestOAuthSignIn(w http.ResponseWriter, r *http.Request, oauthProvider string, clientId string) error {
	if oauthHandler, ok := s.oauth[oauthProvider]; ok {
		expiration := time.Now().Add(5 * time.Minute)
		cookie := http.Cookie{Name: "clientId", Value: clientId, Expires: expiration}
		http.SetCookie(w, &cookie)

		oauthHandler.RedirectToLogin(w, r)
	}

	return nil
}

func (s Identity) FinishOAuthSignIn(w http.ResponseWriter, r *http.Request, oauthProvider string, clientId string) error {
	oauth, ok := s.oauth[oauthProvider]

	if !ok {
		log.Println("Unknown OauthProvider: " + oauthProvider)
		return nil
	}

	oauth.GetUserInfo(w, r)

	jwt := s.jwtAuthority.GenerateJWT(clientId, authorizationCode)
	s.pubsub.Publish(clientId, jwt)

	return nil
}

func (s Identity) SubscribeClient(w http.ResponseWriter, r *http.Request, clientID string) error {

	clientChannel := channel.NewWebSocketChannel(w, r)

	go func() {
		defer clientChannel.Disconnect()
		onJwtReceive := make(chan string)

		subscription := s.pubsub.Subscribe(clientID, func(data interface{}) {
			onJwtReceive <- fmt.Sprint(data) // cast to string
		})
		defer subscription.Unsubscribe()

		select {
		case jwt := <- onJwtReceive:
			// Send JWT signed by Identity service
			clientChannel.SendMessage(jwt)
		case <- time.After(time.Minute * 5):
		}
	}()

	return nil
}

func (s Identity) ValidateAuthToken(authToken string) error {
	panic("not implemented")
}

func NewIdentity(oauthProviders []oauth.OAuth, idGenerator idgen.IDGenerator, userDao dao.User, externalUserDao dao.ExternalUser, jwtAuthority security.JWTAuthority, caesarCipher security.CaesarCipher) Identity {
	oauth := make(map[string] oauth.OAuth)
	for _, provider := range oauthProviders {
		oauth[provider.GetName()] = provider
	}

	return Identity{
		idGenerator:     idGenerator,
		userDao:         userDao,
		externalUserDao: externalUserDao,
		jwtAuthority:    jwtAuthority,
		caesarCipher:    caesarCipher,
		oauth:		     oauth,
		pubsub:    		 pubsub.NewChannelPubSub(),
	}
}

var _ Authentication = (*Identity)(nil)