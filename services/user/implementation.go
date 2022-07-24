package user

import (
	"github.com/google/uuid"
	userModel "github.com/rahmaniali-ir/todo-server/models/user"
)

type iUser struct {
	model userModel.IUser
}

var _ IUser = &iUser{}

func NewService(model userModel.IUser) *iUser {
	return &iUser{
		model: model,
	}
}

func (u *iUser) Get(uid string) (*userModel.PublicUser, error) {
	user, err := u.model.Get(uid)

	return userModel.GetPublicUser(user), err
}

func (u *iUser) Add(name string, email string, password string) (*userModel.User, error) {
	uid := uuid.NewString()
	newUser := &userModel.User{
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
