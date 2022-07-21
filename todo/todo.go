package todo

import (
	"github.com/google/uuid"
)

type Status int

const (
	Status_Undone Status = iota
	Status_In_Progress
	Status_Done
)

type Todo struct {
	Id string `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
  Status Status `json:"status"`
}

type Collection struct {
	Todos map[string]Todo
}

func (c *Collection) ToArray() []Todo {
	todoArray := []Todo{}
	
	for _, todo := range c.Todos {
		todoArray = append(todoArray, todo)
	}

	return todoArray
}

func (c *Collection) AddTodo(todo Todo) Todo {
	todo.Id = uuid.NewString()
	
	c.Todos[todo.Id] = todo

	return todo
}

func (c *Collection) DeleteTodo(uid string) {
	delete(c.Todos, uid)
}

func (c *Collection) ToggleTodo(uid string) Todo {
	todo := c.Todos[uid]

	if todo.Status == Status_Undone {
		todo.Status = Status_Done
	} else if todo.Status == Status_Done {
		todo.Status = Status_Undone
	}

	c.Todos[uid] = todo

	return todo
}
