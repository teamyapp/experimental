package oauth

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/entity"
)

type Facebook struct {
}

func (f Facebook) GetName() string {
	panic("implement me")
}

func (f Facebook) GetSignInURL(clientId string) string {
	panic("implement me")
}

func (f Facebook) GetUserInfo(authorizationCode string) (entity.ExternalUserInfo, error) {
	panic("implement me")
}

func NewFacebook() Facebook {
	return Facebook{}
}

var _ OAuth = (*Facebook)(nil)