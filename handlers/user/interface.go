package user

import "github.com/rahmaniali-ir/todo-server/internal/http"

type IHandler interface {
	Post(*http.GenericRequest) (interface{}, error)
}
