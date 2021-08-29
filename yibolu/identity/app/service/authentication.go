package service

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/oauth"
	"github.com/teamyapp/experimental/yibolu/identity/app/security"

	"github.com/teamyapp/experimental/yibolu/identity/app/dao"
	"github.com/teamyapp/experimental/yibolu/identity/app/idgen"
)

type Authentication interface {
	RequestOAuthSignIn(oauthProvider string, clientId string) error
	FinishOAuthSignIn(oauthProvider string, clientId string, authorizationCode string) error
	ValidateAuthToken(authToken string) error
}

type Identity struct {
	idGenerator     idgen.IDGenerator
	userDao         dao.User
	externalUserDao dao.ExternalUser
	jwtAuthority    security.JWTAuthority
	caesarCipher    security.CaesarCipher
	facebookOAuth   oauth.OAuth
	googleOAuth     oauth.OAuth
	githubOAuth     oauth.OAuth
}

func (s Identity) RequestOAuthSignIn(oauthProvider string, clientId string) error {
	panic("not implemented")
}

func (s Identity) FinishOAuthSignIn(oauthProvider string, clientId string, authorizationCode string) error {
	panic("not implemented")
}

func (s Identity) ValidateAuthToken(authToken string) error {
	panic("not implemented")
}

func NewIdentity(idGenerator idgen.IDGenerator, userDao dao.User, externalUserDao dao.ExternalUser, jwtAuthority security.JWTAuthority, caesarCipher security.CaesarCipher) Identity {
	return Identity{
		idGenerator:     idGenerator,
		userDao:         userDao,
		externalUserDao: externalUserDao,
		jwtAuthority:    jwtAuthority,
		caesarCipher:    caesarCipher,
		facebookOAuth:   oauth.NewFacebook(),
		googleOAuth:     oauth.NewGoogle(),
		githubOAuth:     oauth.NewGithub(),
	}
}
