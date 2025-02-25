package middleware

import (
	"PowerKey/config"
	"PowerKey/server/resp"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func BaseAuth(c *fiber.Ctx) error {
	token := strings.TrimSpace(c.Query("token"))
	if token == "" {
		return c.JSON(resp.Err(resp.TokenNotFound))
	}
	if token != config.App.Server.Token {
		return c.JSON(resp.Err(resp.Unauthorized))
	}
	return c.Next()
}

func ParamCheck(c *fiber.Ctx) error {
	deviceName := c.Query("name")
	if deviceName == "" {
		return c.JSON(resp.Err(resp.InvalidQueryParam))
	}
	device, exists := config.App.Devices[deviceName]
	if !exists {
		return c.JSON(resp.Err(resp.DeviceNotFound))
	}
	c.Locals("device", device)
	return c.Next()
}
