package server

import (
	"PowerKey/server/api"
	"PowerKey/server/middleware"
	"github.com/gofiber/fiber/v2"
)

func Init() {
	app := fiber.New()

	v1 := app.Group("/api/action", middleware.BaseAuth, middleware.ParamCheck)
	v1.Get("/wol", api.ExecuteWol)
	v1.Get("/shutdown", api.ExecuteShutdown)
	v1.Get("/ping", api.ExecutePing)
	app.Listen(":3000")
}
