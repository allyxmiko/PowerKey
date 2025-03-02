package api

import (
	"PowerKey/config"
	"PowerKey/model"
	"PowerKey/server/resp"
	"PowerKey/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleLogin(c *fiber.Ctx) error {
	var loginDto model.User
	if err := c.BodyParser(&loginDto); err != nil {
		return c.JSON(resp.Err(resp.InvalidQueryParam))
	}
	user, err := userService.FindUserByUsername(config.Username)
	if err != nil {
		return c.JSON(resp.Err(resp.UserNotFound))
	}
	if !utils.ComparePasswords(user.Password, loginDto.Password) {
		return c.JSON(resp.Err(resp.PasswordNotMatch))
	}
	token := utils.CreateJwt(config.Username)
	return c.JSON(resp.Data(model.User{Token: token}))
}

func HandleUpdatePassword(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(resp.Err(resp.InvalidQueryParam))
	}

	if _, err := userService.FindUserByUsername(config.Username); err != nil {
		return c.JSON(resp.Err(resp.UserNotFound))
	}
	if err := userService.UpdatePassword(config.Username, user.Password); err != nil {
		return c.JSON(resp.Err(resp.UpdatePasswordFailed))
	}
	return c.JSON(resp.Ok("修改密码成功！"))
}

func HandleUpdateToken(c *fiber.Ctx) error {
	if _, err := userService.FindUserByUsername(config.Username); err != nil {
		return c.JSON(resp.Err(resp.UserNotFound))
	}
	token := utils.RandomString(32)
	if err := userService.UpdateToken(config.Username, token); err != nil {
		return c.JSON(resp.Err(resp.UpdateTokenFailed))
	}
	return c.JSON(resp.Data(model.User{
		Token: token,
	}))
}

func HandleVerifyToken(c *fiber.Ctx) error {
	return c.JSON(resp.Ok("token验证成功！"))
}
