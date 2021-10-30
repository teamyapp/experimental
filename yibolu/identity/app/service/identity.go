package service

import (
	"errors"
	"github.com/teamyapp/experimental/yibolu/identity/app/channel"
	"github.com/teamyapp/experimental/yibolu/identity/app/dao"
	"github.com/teamyapp/experimental/yibolu/identity/app/entity"
	"github.com/teamyapp/experimental/yibolu/identity/app/idgen"
	"github.com/teamyapp/experimental/yibolu/identity/app/oauth"
	"github.com/teamyapp/experimental/yibolu/identity/app/pubsub"
	"github.com/teamyapp/experimental/yibolu/identity/app/security"
	"time"
)

var defaultTimeOut = time.Minute * 5

type Identity struct {
	idGenerator     idgen.IDGenerator
	userDao         dao.User
	externalUserDao dao.ExternalUser
	jwtAuthority    security.JWTAuthority
	caesarCipher    security.CaesarCipher
	oauthProviders  map[string] oauth.OAuth
	pubsub          pubsub.PubSub
	StateManager	oauth.StateManager
}

func (i Identity) RequestOAuthSignInURL(oauthProvider string, clientId string) (string, error) {
	oauthHandler, ok := i.oauthProviders[oauthProvider]
	if !ok {
		return "", errors.New("invalid oauthProviders provider")
	}
	return oauthHandler.GetSignInURL(clientId), nil
}

func (i Identity) FinishOAuthSignIn(authorizationCode string, state entity.OAuthState) error {
	clientID, oauthProvider := state.ClientID, state.OAuthProvider
	oauth, ok := i.oauthProviders[oauthProvider]
	if !ok {
		return errors.New("Unknown OauthProvider: " + oauthProvider)
	}
	externalUserInfo, _ := oauth.GetUserInfo(authorizationCode)
	userInfo := i.getInternalUserInfo(externalUserInfo)

	jwt, _ := i.jwtAuthority.GenerateJWT(userInfo)
	i.pubsub.Publish(clientID, jwt)
	return nil
}

func (s Identity) ClientSubscribe(channel channel.Channel, clientID string) error {
	go func() {
		defer channel.Disconnect()
		onJwtReceive := make(chan string)

		subscription := s.pubsub.Subscribe(clientID, func(data interface{}) {
			jwtToken := data.(string)
			onJwtReceive <- jwtToken
		})
		defer subscription.Unsubscribe()

		select {
		case jwtToken := <- onJwtReceive:
			// Send JWT signed by Identity service
			channel.SendMessage(jwtToken)
		case <- time.After(defaultTimeOut):
		}
	}()
	return nil
}

func (i Identity) ValidateAuthToken(authToken string) error {
	panic("not implemented")
}

// TODO: This function will get the internal user info for the external user
//It will also create a new user when user is not registered yet
func (i Identity) getInternalUserInfo(externalUserInfo entity.ExternalUserInfo) entity.UserInfo {
	panic("not implemented")
}

func NewIdentity(oauthProviders []oauth.OAuth, idGenerator idgen.IDGenerator,
	userDao dao.User, externalUserDao dao.ExternalUser, jwtAuthority security.JWTAuthority,
	caesarCipher security.CaesarCipher, sub pubsub.PubSub, stateManager	oauth.StateManager) Identity {
	oauth := make(map[string] oauth.OAuth)
	for _, provider := range oauthProviders {
		oauth[provider.GetName()] = provider
	}

	return Identity{
		idGenerator:     	idGenerator,
		userDao:         	userDao,
		externalUserDao: 	externalUserDao,
		jwtAuthority:    	jwtAuthority,
		caesarCipher:    	caesarCipher,
		oauthProviders:  	oauth,
		pubsub:          	sub,
		StateManager:		stateManager,
	}
}
