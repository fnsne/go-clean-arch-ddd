package application

import "golang.org/x/crypto/bcrypt"

type BcryptHashTool struct {
}

func (b *BcryptHashTool) Hash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPassword), err
}

func NewBcryptHashTool() *BcryptHashTool {
	return &BcryptHashTool{}
}
