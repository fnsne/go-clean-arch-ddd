package service

import (
	"go-clean-arch-ddd/account/entity"
	"go-clean-arch-ddd/account/usecase/interface/in/register"
	"go-clean-arch-ddd/account/usecase/interface/out"
)

func (s *UserRegisterService) Execute(input register.UserRegisterInput) (register.UserRegisterOutput, error) {
	id := s.userRepository.NextID()
	user := entity.NewUser(id, input.Email, input.Password, input.Name)
	err := s.userRepository.Save(user)
	if err != nil {
		return register.UserRegisterOutput{}, err
	}
	return register.UserRegisterOutput{ID: user.ID}, nil
}

type UserRegisterService struct {
	userRepository out.UserRepository
}

func NewUserRegisterService(userRepository out.UserRepository) *UserRegisterService {
	return &UserRegisterService{userRepository: userRepository}
}
