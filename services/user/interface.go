package user

import (
	"github.com/rahmaniali-ir/todo-server/models/user"
)

type IUser interface {
	Add(name string, email string, password string) (*user.User, error)
}
