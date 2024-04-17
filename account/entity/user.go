package entity

import "errors"

type User struct {
	ID       string
	Email    string
	Password string
	Name     string
}

func (u *User) ChangePassword(password string) {
	u.Password = password
}

func (u *User) Rename(name string) {
	u.Name = name
}

func NewUser(id string, email string, password string, name string) User {
	return User{
		ID:       id,
		Email:    email,
		Password: password,
		Name:     name,
	}
}

var ErrUserNotFound = errors.New("user not found")
