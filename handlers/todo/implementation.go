package todo

import (
	"net/http"

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

func (h *handler) ListAll(req *http.Request) (interface{}, error) {
	return h.service.GetAll()
}
