package todo

import (
	model "github.com/rahmaniali-ir/todo-server/models/todo"
)

type ITodo interface {
	GetAll() ([]model.Todo, error)
}
