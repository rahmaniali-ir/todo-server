package user

type User struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type PublicUser struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type IUser interface {
	Get(uid string) (*User, error)
	Add(user *User) error
}
