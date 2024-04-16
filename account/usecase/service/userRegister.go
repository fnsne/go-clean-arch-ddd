package service

import (
	"go-clean-arch-ddd/account/entity"
	"go-clean-arch-ddd/account/usecase/interface/in/register"
	"go-clean-arch-ddd/account/usecase/interface/out/mocks"
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
	userRepository *mocks.MockUserRepository
}

func NewUserRegisterService(userRepository *mocks.MockUserRepository) *UserRegisterService {
	return &UserRegisterService{userRepository: userRepository}
}
