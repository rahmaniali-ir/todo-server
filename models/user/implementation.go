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

func GetPublicUser(user *User) *PublicUser {
	return &PublicUser{
		Uid: user.Uid,
		Name: user.Name,
		Email: user.Email,
	}
}

func (u *iUser) Get(uid string) (*User, error) {
	userBytes, err := u.db.Get([]byte("user#" + uid), nil)
	if err != nil {
		return nil, err
	}

	user := &User{}
	reader := bytes.NewReader(userBytes)
	err = gob.NewDecoder(reader).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *iUser) Add(user *User) error {
	var userBytes bytes.Buffer
	err := gob.NewEncoder(&userBytes).Encode(user)
	if err != nil {
		return err
	}
	
	return u.db.Put([]byte("user#" + user.Uid), userBytes.Bytes(), nil)
}
