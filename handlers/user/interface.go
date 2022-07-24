package user

import "github.com/rahmaniali-ir/todo-server/internal/http"

type IHandler interface {
	GetProfile(*http.GenericRequest) (interface{}, error)
	SignUp(*http.GenericRequest) (interface{}, error)
}
