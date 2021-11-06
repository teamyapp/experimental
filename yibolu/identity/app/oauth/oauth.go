package oauth

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/entity"
	"net/http"
)

type OAuth interface {
	GetName() string
	GetSignInURL(stateID string) string
	GetUserInfo(authorizationCode string) (entity.ExternalUserInfo, error)
	GetStateID(request * http.Request) (string, error)
	GetAuthorizationCode(request * http.Request) (string, error)
}
