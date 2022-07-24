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

type UserWithToken struct{
	Token string `json:"token"`
	User PublicUser `json:"user"`
}

type Credentials struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type IUser interface {
	Get(uid string) (*User, error)
	GetByCredentials(email string, password string) (*User, error)
	Add(user *User) error
}
