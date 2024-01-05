package user

import "fmt"

// User is a entity in Domain Driven Design
type User struct {
	ID           string
	Name         string
	Email        string
	Status       string
	HashPassword string
}

func NewUser(name string, email string, password string) *User {
	//todo 未來改成用hash的方式
	hashPassword := password
	return &User{
		Name:         name,
		Email:        email,
		HashPassword: hashPassword,
	}
}

func (u *User) PasswordValidate(password string) bool {
	//todo 未來改成用hash的方式
	return u.HashPassword == password
}

func (u *User) ResetPassword(password string) error {
	//todo 未來改成用hash的方式
	u.HashPassword = password
	if u.HashPassword == password {
		return fmt.Errorf("password can not be the same as the old one")
	}
	return nil
}
