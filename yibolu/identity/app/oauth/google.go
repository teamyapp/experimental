package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/teamyapp/experimental/yibolu/identity/app/entity"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	// TODO: randomize it
	oauthStateString = "pseudo-random"

	googleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
)

type Google struct {
	googleOauthConfig *oauth2.Config
}

func NewGoogle() Google {
	return Google{
		googleOauthConfig: &oauth2.Config{
			RedirectURL:  "http://localhost:8080/callback",
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (g Google) GetLoginURL(clientId string) string {
	// TODO: Create util function to construct correct state by clientId
	url := g.googleOauthConfig.AuthCodeURL(clientId)
	return url
}

func (g Google) GetUserInfo(authorizationCode string) (entity.ExternalUserInfo, error) {
	token, err := g.googleOauthConfig.Exchange(oauth2.NoContext, authorizationCode)

	userinfo := entity.ExternalUserInfo{}

	if err != nil {
		return userinfo, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get(googleUserInfoURL + token.AccessToken)
	if err != nil {
		return userinfo, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return userinfo, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	// TODO: Need to figure out the contents
	json.Unmarshal(contents, &userinfo)

	return userinfo, nil
}



