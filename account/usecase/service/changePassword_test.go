package service_test

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go-clean-arch-ddd/account/entity"
	chagnepassword "go-clean-arch-ddd/account/usecase/interface/in/changePassword"
	"go-clean-arch-ddd/account/usecase/interface/out/mocks"
	"go-clean-arch-ddd/account/usecase/service"
	"testing"
)

func TestChangePasswordTestSuite(t *testing.T) {
	suite.Run(t, new(ChangePasswordTestSuite))
}

type ChangePasswordTestSuite struct {
	suite.Suite
}

func (suite *ChangePasswordTestSuite) SetupTest() {
}

func (suite *ChangePasswordTestSuite) Test_changePassword() {
	userRepository := mocks.NewMockUserRepository(suite.T())
	changePasswordService := service.NewUserChangePasswordService(userRepository)
	userRepository.EXPECT().FindByID("newID1").Return(entity.User{
		ID:       "newID1",
		Email:    "user1@email.123",
		Password: "oldPassword123",
		Name:     "user1",
	}, nil)
	userRepository.EXPECT().Save(entity.User{
		ID:       "newID1",
		Email:    "user1@email.123",
		Password: "newPassword",
		Name:     "user1",
	}).Return(nil)
	err := changePasswordService.Execute(chagnepassword.Input{
		ID:       "newID1",
		Password: "newPassword",
	})

	require.NoError(suite.T(), err)
}
