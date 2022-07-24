package todo

import (
	"bytes"
	"encoding/json"

	internalHttp "github.com/rahmaniali-ir/todo-server/internal/http"
	model "github.com/rahmaniali-ir/todo-server/models/todo"
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

func (h *handler) Add(req *internalHttp.GenericRequest) (interface{}, error) {
	todo := model.Todo{}
	reader := bytes.NewReader(req.Body)
	err := json.NewDecoder(reader).Decode(&todo)
	if err != nil {
		return nil, err
	}

	return h.service.AddTodo(todo.Title, todo.Body, todo.Status, req.Session.Uid)
}

func (h *handler) Delete(req *internalHttp.GenericRequest) (interface{}, error) {
	return nil, h.service.DeleteTodo(req.QueryParams["uid"][0])
}
