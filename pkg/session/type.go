package session

import (
	"bytes"
	"encoding/gob"

	"github.com/google/uuid"
	"github.com/rahmaniali-ir/todo-server/internal/models/user"
	"github.com/syndtr/goleveldb/leveldb"
)

type session struct {
	db *leveldb.DB
}

var _ ISession = &session{}

func New(db *leveldb.DB) ISession {
	return &session{
		db: db,
	}
}

func Init(db *leveldb.DB) {
	sessionManager := New(db)
	Default = sessionManager
}

func (s *session) GetByToken(token string) (*TempUserSession, error) {
	uidBytes, err := s.db.Get([]byte("token#" + token), nil)
	if err != nil {
		return nil, err
	}

	uid := string(uidBytes)
	userBytes, err := s.db.Get([]byte("user#" + uid), nil)
	if err != nil {
		return nil, err
	}

	dbUser := &user.User{}
	reader := bytes.NewReader(userBytes)
	err = gob.NewDecoder(reader).Decode(dbUser)
	if err != nil {
		return nil, err
	}

	tempUserSession := &TempUserSession{
		Uid: dbUser.Uid,
		Name: dbUser.Name,
		Email: dbUser.Email,
		Token: token,
	}

	return tempUserSession, nil
}

func (s *session) SetSession(uid string) (string, error) {
	token := uuid.NewString()
	err := s.db.Put([]byte("token#" + token), []byte(uid), nil)

	if err != nil {
		return "", nil
	}

	return token, nil
}

func (s *session) UnsetSession(token string) error {
	return s.db.Delete([]byte("token#" + token), nil)
}
