package todo

import (
	internalHttp "github.com/rahmaniali-ir/todo-server/internal/http"
)

type IHandler interface {
	GetAll(*internalHttp.GenericRequest) (interface{}, error)
	Add(*internalHttp.GenericRequest) (interface{}, error)
}
