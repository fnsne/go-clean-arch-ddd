package application

import (
	"go-clean-arch-ddd/domain/user"
)

type UserService struct {
	userRepo    user.Repository
	userFactory user.UserFactory
}

func (s *UserService) Register(name string, email string, password string) (string, error) {
	u, err := s.userFactory(name, email, password)
	if err != nil {
		return "", err
	}
	id, err := s.userRepo.Create(u)
	if err != nil {
		return "", err
	}
	return id, nil
}

func NewUserService(repository user.Repository, factory user.UserFactory) *UserService {
	return &UserService{
		userRepo:    repository,
		userFactory: factory,
	}
}
