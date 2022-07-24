package user

import (
	"github.com/rahmaniali-ir/todo-server/models/user"
)

type IUser interface {
	Get(uid string) (*user.PublicUser, error)
	Add(name string, email string, password string) (*user.User, error)
}
