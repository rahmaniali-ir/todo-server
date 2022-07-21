package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

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
	filepath string
}

func (c *Collection) openDB() {
	file, err := ioutil.ReadFile(c.filepath)

	if err != nil {
		return
	}

	todos := []Todo{}
	reader := bytes.NewReader(file)
	err = json.NewDecoder(reader).Decode(&todos)

	c.Todos = make(map[string]Todo)
	for _, todo := range todos {
		c.Todos[todo.Id] = todo
	}
}

func (c *Collection) saveToDB() {
	contents, err := json.Marshal(c.ToArray())

	if err != nil {
		return
	}

	file, err := os.Create(c.filepath)
	if err != nil {
		fmt.Println("Error: Couldn't save to database!")
		return
	}

	_, err = io.WriteString(file, string(contents))
	if err != nil {
		fmt.Println("Error: Couldn't save to database!")
		return
	}
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

	c.saveToDB()

	return todo
}

func (c *Collection) DeleteTodo(uid string) {
	delete(c.Todos, uid)

	c.saveToDB()
}

func (c *Collection) ToggleTodo(uid string) Todo {
	todo := c.Todos[uid]

	if todo.Status == Status_Undone {
		todo.Status = Status_Done
	} else if todo.Status == Status_Done {
		todo.Status = Status_Undone
	}

	c.Todos[uid] = todo

	c.saveToDB()

	return todo
}

func NewCollection(filepath string) Collection {
	collection := Collection{
		Todos: make(map[string]Todo),
		filepath: filepath,
	}

	collection.openDB()

	return collection
}
