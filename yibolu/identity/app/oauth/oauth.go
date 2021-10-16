package oauth

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/entity"
)

type OAuth interface {
	GetName() string
	GetSignInURL(clientId string) string
	GetUserInfo(authorizationCode string) (entity.ExternalUserInfo, error)
}
