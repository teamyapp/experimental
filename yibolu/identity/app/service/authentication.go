package service

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/dao"
	"github.com/teamyapp/experimental/yibolu/identity/app/idgen"
	"github.com/teamyapp/experimental/yibolu/identity/app/oauth"
	"github.com/teamyapp/experimental/yibolu/identity/app/queue"
	"github.com/teamyapp/experimental/yibolu/identity/app/security"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var timeWait = time.Second * 300 // 5 minutes

type Authentication interface {
	RequestOAuthSignIn(w http.ResponseWriter, r *http.Request, oauthProvider string, clientId string) error
	FinishOAuthSignIn(w http.ResponseWriter, r *http.Request, oauthProvider string, clientId string, authorizationCode string) error
	ValidateAuthToken(authToken string) error
}

type Identity struct {
	idGenerator     idgen.IDGenerator
	userDao         dao.User
	externalUserDao dao.ExternalUser
	jwtAuthority    security.JWTAuthority
	caesarCipher    security.CaesarCipher
	oauth           map[string] oauth.OAuth
	queue    		queue.MessageQueue
}

func (s Identity) RequestOAuthSignIn(w http.ResponseWriter, r *http.Request, oauthProvider string, clientId string) error {
	panic("not implemented")
}

func (s Identity) FinishOAuthSignIn(oauthProvider string, clientId string, authorizationCode string) error {
	panic("not implemented")
}

func (s Identity) SubscribeClient(w http.ResponseWriter, r *http.Request, clientID string) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	if err != nil {
		log.Println(err)
		return err
	}

	err = s.queue.Subscribe(clientID)
	if err != nil {
		log.Fatal("SubscribeClient: ", err)
		return err
	}

	go s.queue.GetJWT(clientID, func(jwt string) {
		conn.SetWriteDeadline(time.Now().Add(timeWait))
		conn.NextWriter(websocket.TextMessage)

		w, err := conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Fatal(err)
			return
		}

		w.Write([]byte(jwt))

		if err = w.Close(); err != nil {
			log.Fatal(err)
			return
		}
	})

	return err
}

func (s Identity) ValidateAuthToken(authToken string) error {
	panic("not implemented")
}

func NewIdentity(oauthProviders []oauth.OAuth, idGenerator idgen.IDGenerator, userDao dao.User, externalUserDao dao.ExternalUser, jwtAuthority security.JWTAuthority, caesarCipher security.CaesarCipher, queue queue.MessageQueue) Identity {
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
		queue:    		 queue,
	}
}