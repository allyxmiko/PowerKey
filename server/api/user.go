package api

import (
	"PowerKey/config"
	"PowerKey/model/dto"
	"PowerKey/model/vo"
	"PowerKey/server/resp"
	"PowerKey/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleLogin(c *fiber.Ctx) error {
	var loginDto dto.LoginDto
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
	return c.JSON(resp.Data(vo.LoginVo{Token: token}))
}

func HandleUpdatePassword(c *fiber.Ctx) error {
	var loginDto dto.LoginDto
	if err := c.BodyParser(&loginDto); err != nil {
		return c.JSON(resp.Err(resp.InvalidQueryParam))
	}

	if _, err := userService.FindUserByUsername(config.Username); err != nil {
		return c.JSON(resp.Err(resp.UserNotFound))
	}
	if err := userService.UpdatePassword(config.Username, loginDto.Password); err != nil {
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
	return c.JSON(resp.Data(vo.LoginVo{
		Token: token,
	}))
}
