package service

import (
	chagnepassword "go-clean-arch-ddd/account/usecase/interface/in/changePassword"
	"go-clean-arch-ddd/account/usecase/interface/out"
)

type UserChangePasswordService struct {
	userRepository out.UserRepository
}

func NewUserChangePasswordService(userRepository out.UserRepository) *UserChangePasswordService {
	return &UserChangePasswordService{userRepository: userRepository}
}

func (s *UserChangePasswordService) Execute(input chagnepassword.UserChangePasswordInput) error {
	user, err := s.userRepository.FindByID(input.ID)
	if err != nil {
		return err
	}
	user.ChangePassword(input.Password)
	err = s.userRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}
