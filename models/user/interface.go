package user

type User struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type IUser interface {
	Add(user *User) error
}
