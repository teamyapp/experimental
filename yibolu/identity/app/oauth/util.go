package oauth

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/entity"
)

// GetOAuthCallbackState TODO: implement this function, and maybe integrate with redis to get the clientId
func GetOAuthCallbackState(state string) entity.State {
	return entity.State{}
}