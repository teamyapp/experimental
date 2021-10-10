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

var timeWait = time.Minute * 5 // 5 minutes


type Identity struct {
	idGenerator     idgen.IDGenerator
	userDao         dao.User
	externalUserDao dao.ExternalUser
	jwtAuthority    security.JWTAuthority
	caesarCipher    security.CaesarCipher
	oauth           map[string] oauth.OAuth
	pubsub    		pubsub.PubSub
}

func (s Identity) RequestOAuthSignInURL(oauthProvider string, clientId string) (string, error) {
	oauthHandler, ok := s.oauth[oauthProvider]
	if !ok {
		return "", errors.New("invalid oauth provider")
	}

	url := oauthHandler.GetLoginURL(clientId)

	return url, nil
}

func (i Identity) FinishOAuthSignIn(authorizationCode string, oauthProvider string, clientId string) error {
	oauth, ok := i.oauth[oauthProvider]

	if !ok {
		return errors.New("Unknown OauthProvider: " + oauthProvider)
	}

	externalUserInfo, _ := oauth.GetUserInfo(authorizationCode)
	userInfo := i.getInternalUserInfo(externalUserInfo)

	jwt := i.jwtAuthority.GenerateJWT(clientId, userInfo)
	i.pubsub.Publish(clientId, jwt)

	return nil
}

func (s Identity) SubscribeClient(channel channel.Channel, clientID string) error {

	go func() {
		defer channel.Disconnect()
		onJwtReceive := make(chan string)

		subscription := s.pubsub.Subscribe(clientID, func(data interface{}) {
			jwtToken := data.(string)
			onJwtReceive <- jwtToken
		})
		defer subscription.Unsubscribe()

		select {
		case jwt := <- onJwtReceive:
			// Send JWT signed by Identity service
			channel.SendMessage(jwt)
		case <- time.After(time.Minute * 5):
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
	caesarCipher security.CaesarCipher, sub pubsub.PubSub) Identity {
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
		pubsub:    		 sub,
	}
}
