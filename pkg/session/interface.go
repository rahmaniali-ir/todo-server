package session

type TempUserSession struct {
	uid string
	name string
	email string
}

type ISession interface {
	GetByToken(string) (*TempUserSession, error)
}

var Default ISession
