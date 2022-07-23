package session

import (
	"bytes"
	"encoding/gob"

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
