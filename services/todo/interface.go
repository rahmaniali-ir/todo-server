package todo

import (
	model "github.com/rahmaniali-ir/todo-server/models/todo"
)

type ITodo interface {
	GetTodo(uid string) (*model.Todo, error)
	GetUserTodos(userUid string) (*[]model.Todo, error)
	AddTodo(title string, body string, status model.Status, userUid string) (*model.Todo, error)
	DeleteTodo(uid string) error
}
