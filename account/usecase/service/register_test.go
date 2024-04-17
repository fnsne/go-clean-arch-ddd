package service_test

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go-clean-arch-ddd/account/entity"
	"go-clean-arch-ddd/account/usecase/interface/in/register"
	"go-clean-arch-ddd/account/usecase/interface/out/mocks"
	"go-clean-arch-ddd/account/usecase/service"
	"testing"
)

func TestUserRegisterSuite(t *testing.T) {
	suite.Run(t, new(UserRegisterSuite))
}

type UserRegisterSuite struct {
	suite.Suite
}

func (suite *UserRegisterSuite) SetupTest() {

}

func (suite *UserRegisterSuite) Test_UserRegisterSuccess() {
	userRepository := mocks.NewMockUserRepository(suite.T())
	useCase := service.NewUserRegisterService(userRepository)
	userRepository.EXPECT().NextID().Return("newID1")
	userRepository.EXPECT().Save(entity.User{
		ID:       "newID1",
		Email:    "user1@email.123",
		Password: "plainPassword123",
		Name:     "123",
	}).Return(nil)
	output, err := useCase.Execute(register.UserRegisterInput{
		Email:    "user1@email.123",
		Password: "plainPassword123",
		Name:     "123",
	})
	require.NoError(suite.T(), err)
	require.NotEmptyf(suite.T(), output.ID, "ID should not be empty")
}
