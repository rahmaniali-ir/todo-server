package session

type TempUserSession struct {
	Uid string
	Name string
	Email string
	Token string
}

type ISession interface {
	GetByToken(token string) (*TempUserSession, error)
	SetSession(uid string) (token string, err error)
	UnsetSession(token string) error
}

var Default ISession
