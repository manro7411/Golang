package service

type User struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type UserService interface { //port
	SignUp(email string, password string) (*int64, *string, *string, error) //return string pointer ชี้หาตัวแปรตัวนั้น
	SignIn(email string, password string) (*int64, error)
}
