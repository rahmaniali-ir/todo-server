package todo

import (
	"github.com/google/uuid"
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

func (t *iTodo) AddTodo(title string, body string, status model.Status, userUid string) (*model.Todo, error) {
	uid := uuid.NewString()
	todo := model.Todo{
		Uid: uid,
		Title: title,
		Body: body,
		Status: status,
		User_uid: userUid,
	}

	err := t.model.AddTodo(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
