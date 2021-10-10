package security

import "github.com/teamyapp/experimental/yibolu/identity/app/entity"

type JWTAuthority struct {
	signingKey []byte

}

func (j JWTAuthority) GenerateJWT(clientId string, userInfo entity.UserInfo) string {
	return ""
}


