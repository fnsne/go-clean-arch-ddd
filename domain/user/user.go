package user

import "golang.org/x/crypto/bcrypt"

// User is a entity in Domain Driven Design
type User struct {
	ID           string
	Name         string
	Email        string
	Status       string
	HashPassword string
}

func BcryptUserFactory(name string, email string, plainTextPassword string) (*User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:         name,
		Email:        email,
		HashPassword: string(hashPassword),
	}, nil
}
