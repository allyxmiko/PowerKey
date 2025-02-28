package service

import (
	"PowerKey/global"
	"PowerKey/model"
	"PowerKey/utils"
)

type UserService struct{}

type IUserService interface {
	FindByUsername(username string) (model.User, error)
	UpdatePassword(username, password string) error
}

func NewUserService() IUserService {
	return &UserService{}
}

func (u UserService) FindByUsername(username string) (model.User, error) {
	user := model.User{
		Username: username,
	}
	result := global.DB.First(&user)
	return user, result.Error
}

func (u UserService) UpdatePassword(username, password string) error {
	user := model.User{
		Username: username,
		Password: utils.HashPassword(password),
	}
	return global.DB.Save(&user).Error
}
