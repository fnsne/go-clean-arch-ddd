package service_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go-clean-arch-ddd/account/entity"
	"go-clean-arch-ddd/account/usecase/interface/in/getUser"
	"go-clean-arch-ddd/account/usecase/interface/out/mocks"
	"go-clean-arch-ddd/account/usecase/service"
	"testing"
)

func TestGetUserTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserTestSuite))
}

type GetUserTestSuite struct {
	suite.Suite
}

func (suite *GetUserTestSuite) SetupTest() {
}

func (suite *GetUserTestSuite) Test_GetUser() {
	userRepository := mocks.NewMockUserRepository(suite.T())
	useCase := service.NewGetUserUseCase(userRepository)

	userRepository.EXPECT().FindByID("id1").Return(entity.User{
		ID:       "id1",
		Email:    "user1@email.123",
		Password: "userPassword",
		Name:     "123",
	}, nil)
	output, err := useCase.Execute(getuser.Input{
		ID: "id1",
	})
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), "id1", output.ID)
	assert.Equal(suite.T(), "user1@email.123", output.Email)
	assert.Equal(suite.T(), "123", output.Name)
}
