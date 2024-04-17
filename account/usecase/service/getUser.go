package service

import (
	"go-clean-arch-ddd/account/usecase/interface/in/getUser"
	"go-clean-arch-ddd/account/usecase/interface/out"
)

type GetUserUseCase struct {
	userRepository out.UserRepository
}

func (c *GetUserUseCase) Execute(input getuser.Input) (getuser.Output, error) {
	user, err := c.userRepository.FindByID(input.ID)
	if err != nil {
		return getuser.Output{}, err
	}
	return getuser.Output{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func NewGetUserUseCase(userRepository out.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{userRepository: userRepository}
}
