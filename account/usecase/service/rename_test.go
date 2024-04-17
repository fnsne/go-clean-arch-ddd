package service_test

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go-clean-arch-ddd/account/entity"
	"go-clean-arch-ddd/account/usecase/interface/in/rename"
	"go-clean-arch-ddd/account/usecase/interface/out/mocks"
	"go-clean-arch-ddd/account/usecase/service"
	"testing"
)

func TestChangePasswordTestSuite(t *testing.T) {
	suite.Run(t, new(RenameTestSuite))
}

type RenameTestSuite struct {
	suite.Suite
}

func (suite *RenameTestSuite) SetupTest() {
}

func (suite *RenameTestSuite) Test_rename() {
	userRepository := mocks.NewMockUserRepository(suite.T())
	renameService := service.NewUserRenameService(userRepository)
	userRepository.EXPECT().FindByID("newID1").Return(entity.User{
		ID:       "newID1",
		Email:    "user1@email.123",
		Password: "oldPassword123",
		Name:     "user1",
	}, nil)
	userRepository.EXPECT().Save(entity.User{
		ID:       "newID1",
		Email:    "user1@email.123",
		Password: "oldPassword123",
		Name:     "new name 1",
	}).Return(nil)
	err := renameService.Execute(rename.Input{
		ID:   "newID1",
		Name: "new name 1",
	})

	require.NoError(suite.T(), err)
}
