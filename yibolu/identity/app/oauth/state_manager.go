package oauth

import "github.com/teamyapp/experimental/yibolu/identity/app/entity"

type StateManager interface {
	GetOAuthCallbackState(stateID string) entity.OAuthState
	SaveOAuthState(state entity.OAuthState)
}

type CacheStateManager struct {
	cache			map[string]interface{}
}

// GetOAuthCallbackState TODO: implement this function, and maybe integrate with redis to get the clientId
func (c *CacheStateManager) GetOAuthCallbackState(stateID string) entity.OAuthState {
	return entity.OAuthState{}
}

// SaveOAuthState TODO: implement this function, and maybe integrate with redis to get the clientId
func (c *CacheStateManager) SaveOAuthState(entity.OAuthState) error {
	return nil
}