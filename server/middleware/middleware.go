package middleware

import (
	"PowerKey/config"
	"PowerKey/server/resp"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func BaseAuth(c *fiber.Ctx) error {
	if !config.App.Server.Auth {
		return c.Next()
	}
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

	for _, device := range config.App.Devices {
		if device.Name == deviceName {
			c.Locals("device", device)
			return c.Next()
		}
	}
	return c.JSON(resp.Err(resp.DeviceNotFound))
}
