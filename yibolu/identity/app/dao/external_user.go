package dao

import "github.com/teamyapp/experimental/yibolu/identity/app/entity"

type ExternalUser interface {
	GetInternalUserId(externalUserInfo entity.ExternalUserInfo) string
	CreateExternalUser(externalUserInfo entity.ExternalUserInfo)
	UpdateExternalUser(externalUserInfo entity.ExternalUserInfo)
	DeleteExternalUser(externalUserInfo entity.ExternalUserInfo)
}

type ExternalUserSQL struct {
}

var _ ExternalUser = (*ExternalUserSQL)(nil)


func (e ExternalUserSQL) GetInternalUserId(externalUserInfo entity.ExternalUserInfo) string {
	panic("implement me")
}

func (e ExternalUserSQL) CreateExternalUser(externalUserInfo entity.ExternalUserInfo) {
	panic("implement me")
}

func (e ExternalUserSQL) UpdateExternalUser(externalUserInfo entity.ExternalUserInfo) {
	panic("implement me")
}

func (e ExternalUserSQL) DeleteExternalUser(externalUserInfo entity.ExternalUserInfo) {
	panic("implement me")
}

func NewExternalUserSQL() ExternalUserSQL {
	return ExternalUserSQL{}
}

