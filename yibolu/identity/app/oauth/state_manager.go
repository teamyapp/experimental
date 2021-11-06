package oauth

import "github.com/teamyapp/experimental/yibolu/identity/app/entity"

type StateManager interface {
	GetOAuthState(stateID string) (entity.OAuthState, error)
	SaveOAuthState(stateID string, state entity.OAuthState) error
}

type InMemoryStateManager struct {
	cache			map[string]interface{}
}

func (c *InMemoryStateManager) GetOAuthState(stateID string) (entity.OAuthState, error) {
	//TODO: implement this function, and maybe integrate with redis to get the clientId
	return entity.OAuthState{}, nil
}

func (c *InMemoryStateManager) SaveOAuthState(stateID string, state entity.OAuthState) error {
	//TODO: implement this function, and maybe integrate with redis to get the clientId
	return nil
}