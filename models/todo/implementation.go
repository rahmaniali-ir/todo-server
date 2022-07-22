package todo

import (
	"bytes"
	"encoding/gob"

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

func (t *iTodo) GetAll() ([]Todo, error) {
	iter := t.db.NewIterator(util.BytesPrefix([]byte("todo#")), nil)

	todos := []Todo{}
	for iter.Next() {
		todo := Todo{}
		reader := bytes.NewReader(iter.Value())
		err := gob.NewDecoder(reader).Decode(&todo)
		if err != nil {
			return []Todo{}, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
