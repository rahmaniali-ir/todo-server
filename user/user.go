package user

type User struct {
	Uid string `json:"uid"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}
