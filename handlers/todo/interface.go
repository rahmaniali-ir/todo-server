package todo

import (
	internalHttp "github.com/rahmaniali-ir/todo-server/internal/http"
)

type IHandler interface {
	ListAll(*internalHttp.GenericRequest) (interface{}, error)
}
