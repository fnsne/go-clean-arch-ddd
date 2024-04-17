package service

import (
	"go-clean-arch-ddd/account/entity"
	"go-clean-arch-ddd/account/usecase/interface/in/register"
	"go-clean-arch-ddd/account/usecase/interface/out"
)

func (s *UserRegisterService) Execute(input register.Input) (register.Output, error) {
	id := s.userRepository.NextID()
	user := entity.NewUser(id, input.Email, input.Password, input.Name)
	err := s.userRepository.Save(user)
	if err != nil {
		return register.Output{}, err
	}
	return register.Output{ID: user.ID}, nil
}

type UserRegisterService struct {
	userRepository out.UserRepository
}

func NewUserRegisterService(userRepository out.UserRepository) *UserRegisterService {
	return &UserRegisterService{userRepository: userRepository}
}
