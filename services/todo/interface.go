package todo

import (
	model "github.com/rahmaniali-ir/todo-server/models/todo"
)

type ITodo interface {
	GetUserTodos(userUid string) ([]model.Todo, error)
	AddTodo(title string, body string, status model.Status, userUid string) (*model.Todo, error)
	DeleteTodo(uid string) error
}
