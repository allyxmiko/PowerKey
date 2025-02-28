package api

import (
	"PowerKey/model/dto"
	"PowerKey/server/resp"
	"github.com/gofiber/fiber/v2"
)

func HandleLogin(c *fiber.Ctx) error {
	var loginDto dto.LoginDto
	if err := c.BodyParser(&loginDto); err != nil {
		return c.JSON(resp.Err(resp.InvalidQueryParam))
	}

	return c.JSON(resp.Data(loginDto))
}
