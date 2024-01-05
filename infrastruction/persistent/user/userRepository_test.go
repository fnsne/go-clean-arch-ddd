package user_test

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	userdomain "go-clean-arch-ddd/domain/user"
	userinfra "go-clean-arch-ddd/infrastruction/persistent/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

type UserRepositorySuite struct {
	suite.Suite
	Repository *userinfra.GormUserRepository
}

func (suite *UserRepositorySuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	require.NoError(suite.T(), err)
	suite.Repository = userinfra.NewGormUserRepository(db)
}

func (suite *UserRepositorySuite) Test_Create() {
	userID, err := suite.Repository.Create(&userdomain.User{
		ID:           "userID1",
		Name:         "name",
		Email:        "email",
		HashPassword: "password",
		Status:       "status",
	})
	require.NoError(suite.T(), err)
	require.NotEmpty(suite.T(), userID)
}
