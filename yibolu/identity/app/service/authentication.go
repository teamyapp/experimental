package service

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/oauth"
	"github.com/teamyapp/experimental/yibolu/identity/app/security"
	"go/types"

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
	oauth           map[string] oauth.OAuth
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

//NewIdentity([]oauth.OAuth{oauth.NewGoogle(), oauth.NewGithub(), oauth.NewFacebook()})

func oauthSignIn(provider string) {
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).

	config file
	{
		oauth providers : ["google", "github"]
	}

	config := getConfig(provider)

	conf := &oauth2.Config{
		ClientID:     config.Id,
		ClientSecret: config.Secret,
		RedirectURL:  "YOUR_REDIRECT_URL",
		Scopes: []string{
			"https://www.googleapis.com/auth/bigquery",
			"https://www.googleapis.com/auth/blogger",
		},
		Endpoint: google.Endpoint,
	}
	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state")
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Handle the exchange code to initiate a transport.
	tok, err := conf.Exchange(oauth2.NoContext, "authorization-code")
	if err != nil {
	log.Fatal(err)
	}
	client := conf.Client(oauth2.NoContext, tok)
	client.Get("...")
}

func NewIdentity(oauthProviders []oauth.OAuth,idGenerator idgen.IDGenerator, userDao dao.User, externalUserDao dao.ExternalUser, jwtAuthority security.JWTAuthority, caesarCipher security.CaesarCipher) Identity {
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
	}
}