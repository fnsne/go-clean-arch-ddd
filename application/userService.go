package application

import (
	"go-clean-arch-ddd/application/mocks"
	"go-clean-arch-ddd/domain/user"
)

type UserService struct {
	userRepo user.Repository
	hashTool *mocks.MockHashTool
}

func (s *UserService) Register(name string, email string, password string) (string, error) {
	hashPassword, err := s.hashTool.Hash(password)
	if err != nil {
		return "", err
	}
	u := user.User{
		Name:         name,
		Email:        email,
		HashPassword: hashPassword,
	}
	id, err := s.userRepo.Create(&u)
	if err != nil {
		return "", err
	}
	return id, nil
}

func NewUserService(repository user.Repository, hashTool *mocks.MockHashTool) *UserService {
	return &UserService{
		userRepo: repository,
		hashTool: hashTool,
	}
}
