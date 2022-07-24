package session

type TempUserSession struct {
	Uid string
	Name string
	Email string
}

type ISession interface {
	GetByToken(token string) (*TempUserSession, error)
	SetSession(uid string) (token string, err error)
}

var Default ISession
