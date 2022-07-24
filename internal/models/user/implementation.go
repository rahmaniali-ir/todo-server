package user

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
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

func (u *iUser) GetByCredentials(email string, password string) (*User, error) {
	iter := u.db.NewIterator(util.BytesPrefix([]byte("user#")), nil)

	for iter.Next() {
		user := &User{}
		reader := bytes.NewReader(iter.Value())
		err := gob.NewDecoder(reader).Decode(&user)
		if err != nil {
			return nil, err
		}

		if user.Email == email && user.Password == password {
			return user, nil
		}
	}

	return nil, errors.New("Email or password is wrong!")
}

func (u *iUser) Add(user *User) error {
	var userBytes bytes.Buffer
	err := gob.NewEncoder(&userBytes).Encode(user)
	if err != nil {
		return err
	}
	
	return u.db.Put([]byte("user#" + user.Uid), userBytes.Bytes(), nil)
}
