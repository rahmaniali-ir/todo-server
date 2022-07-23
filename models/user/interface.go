package user

type User struct {
	Uid string
	Name string
	Email string
	Password string
}

type IUser interface {
	Add(user *User) error
}
