package service

import (
	"PowerKey/db"
	"PowerKey/model"
	"PowerKey/utils"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

type IUserService interface {
	FindUserByUsername(username string) (model.User, error)
	UpdatePassword(username, password string) error
}

func NewUserService() IUserService {
	return &UserService{
		db: db.DB,
	}
}

func (u UserService) FindUserByUsername(username string) (model.User, error) {
	user := model.User{
		Username: username,
	}
	result := u.db.First(&user)
	return user, result.Error
}

func (u UserService) UpdatePassword(username, password string) error {
	user := model.User{
		Username: username,
		Password: utils.HashPassword(password),
	}
	return u.db.Save(&user).Error
}
