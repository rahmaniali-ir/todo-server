package session

type TempUserSession struct {
	uid string
	name string
	email string
}

type ISession interface {
	GetByToken(token string) (*TempUserSession, error)
	SetSession(uid string) (token string, err error)
}

var Default ISession
