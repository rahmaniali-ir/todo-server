package user

import (
	"bytes"
	"encoding/gob"

	"github.com/syndtr/goleveldb/leveldb"
)

type iUser struct {
	db *leveldb.DB
}

var _ IUser = &iUser{}

func NewModel(db *leveldb.DB) *iUser {
	return &iUser{
		db: db,
	}
}

func (u *iUser) Add(user *User) error {
	var userBytes bytes.Buffer
	err := gob.NewEncoder(&userBytes).Encode(user)
	if err != nil {
		return err
	}
	
	return u.db.Put([]byte("user#" + user.Uid), userBytes.Bytes(), nil)
}
