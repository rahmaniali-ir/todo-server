package todo

import (
	model "github.com/rahmaniali-ir/todo-server/models/todo"
	service "github.com/rahmaniali-ir/todo-server/models/todo"
)

type iTodo struct {
	model service.ITodo
}

var _ ITodo = &iTodo{}

func NewService(todoModel service.ITodo) iTodo {
	t := iTodo{
		model: todoModel,
	}
	
	return t
}

func (t *iTodo) GetUserTodos(userUid string) ([]model.Todo, error) {
	return t.model.GetUserTodos(userUid)
}
