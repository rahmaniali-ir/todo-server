package user

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/rahmaniali-ir/todo-server/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	userDBPath = "./db/user"
	tokenDBPath = "./db/auth-token"
)

type User struct {
	Uid string `json:"uid"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type Collection struct {
	usersDB leveldb.DB
	tokenDB leveldb.DB
}

func (c *Collection) openDB() {
	usersDB, err := leveldb.OpenFile(userDBPath, nil)
	if err != nil {
		panic("Could not open database!")
	}

	tokenDB, err := leveldb.OpenFile(tokenDBPath, nil)
	if err != nil {
		panic("Could not open database!")
	}

	c.usersDB = *usersDB
	c.tokenDB = *tokenDB
}

func (c *Collection) Close() error {
	err := c.usersDB.Close()
	if err != nil {
		return err
	}

	err = c.tokenDB.Close()

	return err
}

func (c *Collection) GetUserByUid(uid string) (User, error) {
	dbUserBytes, err := c.usersDB.Get([]byte(uid), nil)
	if err != nil {
		return User{}, errors.New("Invalid user uid!")
	}

	dbUser := User{}
	reader := bytes.NewReader(dbUserBytes)
	err = gob.NewDecoder(reader).Decode(&dbUser)
	if err == nil {}

	return dbUser, nil
}

func (c *Collection) GetUserUidByToken(token string) (string, error) {
	uidBytes, err := c.tokenDB.Get([]byte(token), nil)
	if err != nil {
		return "", errors.New("Invalid user uid!")
	}

	return string(uidBytes), nil
}

func (c *Collection) GetUserByHeaderToken(r *http.Request) (User, error) {
	token := utils.GetAuthHeaderToken(r)

	uid, err := c.GetUserUidByToken(token)
	if err != nil {
		return User{}, errors.New("Invalid user token!")
	}

	return c.GetUserByUid(uid)
}

func (c *Collection) SearchUsers(selector func (User) bool) ([]User, error) {
	users := []User{}

	iter := c.usersDB.NewIterator(nil, nil)
	for iter.Next() {
		var dbUser User
		reader := bytes.NewReader(iter.Value())
		err := gob.NewDecoder(reader).Decode(&dbUser)

		fmt.Println(dbUser)
		if err != nil {
			return []User{}, err
		}

		if selector(dbUser) {
			users = append(users, dbUser)
		}
	}

	return users, nil
}

func (c *Collection) SearchSingleUser(filter func(User) bool) (User, error) {
	users, err := c.SearchUsers(filter)

	fmt.Println(users, len(users), err)

	if err != nil {
		return User{}, err
	}

	if len(users) == 0 {
		return User{}, errors.New("User not found!")
	}

	return users[0], nil
}

func (c *Collection) AddUser(user User) (string, User, error) {
	uid := uuid.NewString()
	user.Uid = uid

	var userBytes bytes.Buffer
	err := gob.NewEncoder(&userBytes).Encode(user)
	if err != nil {
		return "", User{}, err
	}
	
	err = c.usersDB.Put([]byte(uid), userBytes.Bytes(), nil)
	if err != nil {
		return "", User{}, err
	}

	return uid, user, nil
}

func (c *Collection) SignUserIn(uid string) (string, error) {
	// generate token
	token := uuid.NewString()
	err := c.tokenDB.Put([]byte(token), []byte(uid), nil)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (c *Collection) SignUserOut(token string) error {
	return c.tokenDB.Delete([]byte(token), nil)
}

func NewCollection() Collection {
	c := Collection{}
	c.openDB()

	return c
}
