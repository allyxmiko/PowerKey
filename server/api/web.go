package api

import (
	"PowerKey/config"
	"PowerKey/model/dto"
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

	return c.JSON(resp.Data(loginDto))
}
