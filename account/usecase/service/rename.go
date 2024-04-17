package service

import (
	"go-clean-arch-ddd/account/usecase/interface/in/rename"
	"go-clean-arch-ddd/account/usecase/interface/out"
)

type UserRenameService struct {
	userRepository out.UserRepository
}

func NewUserRenameService(userRepository out.UserRepository) *UserRenameService {
	return &UserRenameService{userRepository: userRepository}
}

func (s *UserRenameService) Execute(input rename.Input) error {
	user, err := s.userRepository.FindByID(input.ID)
	if err != nil {
		return err
	}
	user.Rename(input.Name)
	err = s.userRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}
