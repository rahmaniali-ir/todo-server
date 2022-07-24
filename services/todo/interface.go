package todo

import (
	model "github.com/rahmaniali-ir/todo-server/models/todo"
)

type ITodo interface {
	GetUserTodos(userUid string) ([]model.Todo, error)
}
