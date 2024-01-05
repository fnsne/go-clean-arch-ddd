package application

import (
	"fmt"
	"go-clean-arch-ddd/domain/user"
)

type UserService struct {
	userRepo user.Repository
}

func NewUserService(repository user.Repository) *UserService {
	return &UserService{
		userRepo: repository,
	}
}

func (h *UserService) Register(name string, email string, password string) (string, error) {
	u := user.NewUser(name, email, password)
	id, err := h.userRepo.Create(u)
	return id, err
}

func (h *UserService) Login(email string, password string) (string, error) {
	u, err := h.userRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	validate := u.PasswordValidate(password)
	if !validate {
		return "", fmt.Errorf("password is not correct")
	}
	return u.ID, nil
}
func (h *UserService) PasswordReset(email string, password string) error {
	u, err := h.userRepo.GetByEmail(email)
	if err != nil {
		return err
	}
	return u.ResetPassword(password)
}
