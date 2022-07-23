package user

import (
	"github.com/google/uuid"
	"github.com/rahmaniali-ir/todo-server/models/user"
)

type iUser struct {
	model user.IUser
}

var _ IUser = &iUser{}

func NewService(model user.IUser) *iUser {
	return &iUser{
		model: model,
	}
}

func (u *iUser) Add(name string, email string, password string) (*user.User, error) {
	uid := uuid.NewString()
	newUser := &user.User{
		Uid: uid,
		Name: name,
		Email: email,
		Password: password,
	}

	err := u.model.Add(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
