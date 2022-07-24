package todo

import (
	"bytes"
	"encoding/json"
	"errors"

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
	if req.Session == nil {
		return nil, errors.New("Unauthorized request!")
	}

	todos, err := h.service.GetUserTodos(req.Session.Uid)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (h *handler) Add(req *internalHttp.GenericRequest) (interface{}, error) {
	if req.Session == nil {
		return nil, errors.New("Unauthorized request!")
	}

	todo := model.Todo{}
	reader := bytes.NewReader(req.Body)
	err := json.NewDecoder(reader).Decode(&todo)
	if err != nil {
		return nil, err
	}

	return h.service.AddTodo(todo.Title, todo.Body, todo.Status, req.Session.Uid)
}

func (h *handler) Delete(req *internalHttp.GenericRequest) (interface{}, error) {
	if req.Session == nil {
		return nil, errors.New("Unauthorized request!")
	}

	if !req.QueryParams.Has("uid") {
		return nil, errors.New("Incomplete request!")
	}

	uid := req.QueryParams["uid"][0]
	todo, err := h.service.GetTodo(uid)
	if err != nil {
		return nil, err
	}

	if todo.User_uid != req.Session.Uid {
		return nil, errors.New("Unauthorized request!")
	}

	return nil, h.service.DeleteTodo(uid)
}
