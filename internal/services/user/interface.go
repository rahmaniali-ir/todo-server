package user

import (
	"github.com/rahmaniali-ir/todo-server/internal/models/user"
)

type IUser interface {
	Get(uid string) (*user.PublicUser, error)
	Add(name string, email string, password string) (*user.PublicUser, error)
	GetByCredentials(email string, password string) (*user.PublicUser, error)
}
