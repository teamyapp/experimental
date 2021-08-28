package dao

import "github.com/teamyapp/identity/app/entity"

type User interface {
	GetUserByInternalID(internalID string) entity.User
	CreateUser(user entity.User) error
	UpdateUser(user entity.User) error
	DeleteUser(userID entity.User) error
}

type UserSQL struct {
}

var _ User = (*UserSQL)(nil)

func (u UserSQL) GetUserByInternalID(internalID string) entity.User {
	panic("implement me")
}

func (u UserSQL) CreateUser(user entity.User) error {
	panic("implement me")
}

func (u UserSQL) UpdateUser(user entity.User) error {
	panic("implement me")
}

func (u UserSQL) DeleteUser(userID entity.User) error {
	panic("implement me")
}

func NewUserSQL() UserSQL {
	return UserSQL{}
}


