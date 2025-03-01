package service

import (
	"PowerKey/db"
	"PowerKey/model"
	"PowerKey/utils"
)

type UserService struct {
}

type IUserService interface {
	FindUserByUsername(username string) (model.User, error)
	UpdatePassword(username, password string) error
}

func NewUserService() IUserService {
	return &UserService{}
}

func (u UserService) FindUserByUsername(username string) (model.User, error) {
	user := model.User{
		Username: username,
	}
	result := db.DB.First(&user)
	return user, result.Error
}

func (u UserService) UpdatePassword(username, password string) error {
	user := model.User{
		Username: username,
		Password: utils.HashPassword(password),
	}
	return db.DB.Save(&user).Error
}
