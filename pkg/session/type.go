package session

import (
	"bytes"
	"encoding/gob"

	"github.com/google/uuid"
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
	sessionBytes, err := s.db.Get([]byte("token#" + token), nil)

	if err != nil {
		return &TempUserSession{}, err
	}

	tempUserSession := TempUserSession{}
	reader := bytes.NewReader(sessionBytes)
	err = gob.NewDecoder(reader).Decode(&tempUserSession)

	if err != nil {
		return &TempUserSession{}, err
	}

	return &TempUserSession{}, nil
}

func (s *session) SetSession(uid string) (string, error) {
	token := uuid.NewString()
	err := s.db.Put([]byte("token#" + token), []byte(uid), nil)

	if err != nil {
		return "", nil
	}

	return token, nil
}
