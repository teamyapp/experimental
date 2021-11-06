package oauth

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/entity"
	"net/http"
)

const clientID = "4da1b4a1f09b0ba7a81e"
const clientSecret = "7ebe0c784eaa7f836d373bb0ced17ee0bffda1dd"

type Github struct {

}

func (g Github) GetStateID(request *http.Request) (string, error) {
	panic("implement me")
}

func (g Github) GetAuthorizationCode(request *http.Request) (string, error) {
	panic("implement me")
}

func (g Github) GetName() string {
	panic("implement me")
}

func (g Github) GetSignInURL(clientId string) string {
	panic("implement me")
}

func (g Github) GetUserInfo(authorizationCode string) (entity.ExternalUserInfo, error) {
	panic("implement me")
}

func NewGithub() Github {
	return Github{}
}

var _ OAuth = (*Github)(nil)