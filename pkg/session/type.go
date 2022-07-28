package session

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/rahmaniali-ir/todo-server/internal/models/user"
	"github.com/syndtr/goleveldb/leveldb"
)

type session struct {
	userDB *leveldb.DB
	db *redis.Client
	ctx *context.Context
}

var _ ISession = &session{}

func New(db *redis.Client, ctx *context.Context, userDb *leveldb.DB) ISession {
	return &session{
		db: db,
		ctx: ctx,
		userDB: userDb,
	}
}

func Init(userDb *leveldb.DB) {
	ctx := context.Background()
	redisDB := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

	sessionManager := New(redisDB, &ctx, userDb)
	Default = sessionManager
}

func (s *session) GetByToken(token string) (*TempUserSession, error) {
	uid, err := s.db.Get(*s.ctx, "token#" + token).Result()
	if err != nil {
		return nil, err
	}

	userBytes, err := s.userDB.Get([]byte("user#" + uid), nil)
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
	result := s.db.Set(*s.ctx, "token#" + token, []byte(uid), 24 * 60 * time.Minute)

	if result.Err() != nil {
		return "", nil
	}

	return token, nil
}

func (s *session) UnsetSession(token string) error {
	return s.db.Del(*s.ctx, "token#" + token).Err()
}
