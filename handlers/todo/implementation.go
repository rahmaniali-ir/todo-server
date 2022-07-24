package todo

import (
	internalHttp "github.com/rahmaniali-ir/todo-server/internal/http"
	service "github.com/rahmaniali-ir/todo-server/services/todo"
)

type handler struct {
	service service.ITodo
}

var _ IHandler = &handler{}

func NewHandler(service service.ITodo) IHandler {
	h := &handler{
		service: service,
	}

	return h
}

func (h *handler) GetAll(req *internalHttp.GenericRequest) (interface{}, error) {
	return h.service.GetUserTodos(req.Session.Uid)
}
