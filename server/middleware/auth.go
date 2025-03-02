package middleware

import (
	"PowerKey/server/resp"
	"PowerKey/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var skip = []string{
	"/api/user/login",
}

func hasPath(c *fiber.Ctx) bool {
	for _, s := range skip {
		if c.Path() == s {
			return true
		}
	}
	return false
}

func JwtAuth(c *fiber.Ctx) error {
	if hasPath(c) {
		return c.Next()
	}
	token := c.Get("Authorization")
	if token == "" {
		return c.JSON(resp.Err(resp.TokenNotFound))
	}

	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}

	if _, err := utils.ParseJwt(token); err != nil {
		return c.JSON(resp.Err(resp.Unauthorized))
	}
	return c.Next()
}
