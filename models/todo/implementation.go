package todo

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type iTodo struct {
	db *leveldb.DB
}

var _ ITodo = &iTodo{}

func NewModel(db *leveldb.DB) (ITodo, error) {
	model := &iTodo{}
	model.db = db
	
	return model, nil
}

func (t *iTodo) GetAll() (*[]Todo, error) {
	iter := t.db.NewIterator(util.BytesPrefix([]byte("todo#")), nil)

	todos := []Todo{}
	for iter.Next() {
		todo := Todo{}
		reader := bytes.NewReader(iter.Value())
		err := gob.NewDecoder(reader).Decode(&todo)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return &todos, nil
}

func (t *iTodo) GetTodo(uid string) (*Todo, error) {
	todoBytes, err := t.db.Get([]byte("todo#" + uid), nil)
	if err != nil {
		return nil, errors.New("Data not found!")
	}

	todo := &Todo{}
	reader := bytes.NewReader(todoBytes)
	err = gob.NewDecoder(reader).Decode(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *iTodo) GetUserTodos(userUid string) (*[]Todo, error) {
	iter := t.db.NewIterator(util.BytesPrefix([]byte("todo#")), nil)

	todos := []Todo{}
	for iter.Next() {
		todo := Todo{}
		reader := bytes.NewReader(iter.Value())
		err := gob.NewDecoder(reader).Decode(&todo)
		if err != nil {
			return nil, err
		}

		if todo.User_uid == userUid {
			todos = append(todos, todo)
		}
	}

	return &todos, nil
}

func (t *iTodo) AddTodo(todo *Todo) error {
	var todoBytes bytes.Buffer
	err := gob.NewEncoder(&todoBytes).Encode(todo)
	if err != nil {
		return err
	}

	return t.db.Put([]byte("todo#" + todo.Uid), todoBytes.Bytes(), nil)
}

func (t *iTodo) DeleteTodo(uid string) error {
	err := t.db.Delete([]byte("todo#" + uid), nil)
	if err != nil {
		return errors.New("Data not found!")
	}

	return nil
}
