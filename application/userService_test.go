package application_test

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go-clean-arch-ddd/application"
	"go-clean-arch-ddd/application/mocks"
	userdomain "go-clean-arch-ddd/domain/user"
	usermocks "go-clean-arch-ddd/domain/user/mocks"
	"testing"
)

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

type UserServiceSuite struct {
	suite.Suite
	mockUserRepo *usermocks.MockRepository
	userService  *application.UserService
	mockHashTool *mocks.MockHashTool
}

func (suite *UserServiceSuite) SetupTest() {
	suite.mockUserRepo = usermocks.NewMockRepository(suite.T())
	suite.mockHashTool = mocks.NewMockHashTool(suite.T())
	suite.userService = application.NewUserService(suite.mockUserRepo, suite.mockHashTool)
}

func (suite *UserServiceSuite) Test_Register() {
	suite.mockHashTool.EXPECT().Hash("password").Return("hashPassword", nil)
	suite.mockUserRepo.EXPECT().Create(&userdomain.User{
		Name:         "name",
		Email:        "email",
		HashPassword: "hashPassword",
	}).Return("userID1", nil)
	userID, err := suite.userService.Register("name", "email", "password")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "userID1", userID)
}
