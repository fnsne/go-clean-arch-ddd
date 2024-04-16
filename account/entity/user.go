package entity

type User struct {
	ID       string
	Email    string
	Password string
	Name     string
}

func NewUser(id string, email string, password string, name string) User {
	return User{
		ID:       id,
		Email:    email,
		Password: password,
		Name:     name,
	}
}
