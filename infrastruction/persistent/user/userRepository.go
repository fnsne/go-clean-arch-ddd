package user

import (
	"github.com/lithammer/shortuuid"
	"go-clean-arch-ddd/domain/user"
	"gorm.io/gorm"
)

type User struct {
	ID     string
	Name   string
	Email  string
	Status string
}
type Login struct {
	Email    string
	Password string
	UserID   string
}

func NewGormUserRepository(gormDB *gorm.DB) *GormUserRepository {
	return &GormUserRepository{gormDB: gormDB}
}

type GormUserRepository struct {
	gormDB *gorm.DB
}

func (g *GormUserRepository) GetByEmail(email string) (*user.User, error) {
	var login Login
	err := g.gormDB.Where("email = ?", email).First(&login).Error
	if err != nil {
		return nil, err
	}
	var u user.User
	err = g.gormDB.Where("id = ?", login.UserID).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (g *GormUserRepository) Create(user *user.User) (id string, err error) {
	userID := shortuuid.New()
	tx := g.gormDB.Begin()
	err = tx.Create(&User{
		ID:     userID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}
	err = tx.Create(&Login{
		UserID:   userID,
		Email:    user.Email,
		Password: user.HashPassword,
	}).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return userID, tx.Commit().Error
}

func (g *GormUserRepository) Get(id string) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormUserRepository) Edit(id string, user *user.User) error {
	//TODO implement me
	panic("implement me")
}

func (g *GormUserRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
