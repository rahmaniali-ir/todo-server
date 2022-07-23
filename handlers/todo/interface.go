package todo

import (
	internalHttp "github.com/rahmaniali-ir/todo-server/internal/http"
)

type IHandler interface {
	Get(*internalHttp.GenericRequest) (interface{}, error)
}
