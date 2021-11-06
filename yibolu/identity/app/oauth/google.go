package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/teamyapp/experimental/yibolu/identity/app/config"
	"github.com/teamyapp/experimental/yibolu/identity/app/entity"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	// TODO: randomize it
	oauthStateString = "pseudo-random"
	googleUserInfoURL, _ = url.Parse("https://www.googleapis.com/oauth2/v2/userinfo")
)

type Google struct {
	googleOauthConfig *oauth2.Config
}

func (g Google) GetStateID(request *http.Request) (string, error) {
	stateID := request.FormValue("state")
	if stateID == "" {
		return "", errors.New("")
	}
	return stateID, nil
}

func (g Google) GetAuthorizationCode(request *http.Request) (string, error) {
	panic("implement me")
}

func (g Google) GetName() string {
	return "google"
}

func NewGoogle(config config.Config) Google {
	return Google{
		googleOauthConfig: &oauth2.Config{
			RedirectURL:  "http://localhost:8080/callback",
			ClientID:     config.GoogleClientID,
			ClientSecret: config.GoogleClientSecret,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (g Google) GetSignInURL(stateID string) string {
	// TODO: Create util function to construct correct state by clientId
	url := g.googleOauthConfig.AuthCodeURL(stateID)
	return url
}

func (g Google) GetUserInfo(authorizationCode string) (entity.ExternalUserInfo, error) {
	accessToken, err := g.googleOauthConfig.Exchange(oauth2.NoContext, authorizationCode)

	userInfo := entity.ExternalUserInfo{}

	if err != nil {
		return userInfo, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get(getTokenAccessURL(accessToken.AccessToken))
	if err != nil {
		return userInfo, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return userInfo, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	userInfoMap := make(map[string]string)
	// TODO: Need to figure out the contents
	json.Unmarshal(contents, &userInfoMap)

	return userInfo, nil
}


func getTokenAccessURL(accessToken string) string {
	q := googleUserInfoURL.Query()
	q.Set("access_token", accessToken)
	return googleUserInfoURL.String()
}

var _ OAuth = (*Google)(nil)
