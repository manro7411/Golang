package repository

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

type UserRepository interface {
	CreateUser(email string, password string, secret string) (*User, error) //*as pointer
	CheckUser(Email string) (*User, error)
}
