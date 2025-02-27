package server

import (
	"PowerKey/server/api"
	"PowerKey/server/middleware"
	"PowerKey/server/static"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
)

func Init() {
	app := fiber.New()

	v1 := app.Group("/api")
	action := v1.Group("/action", middleware.BaseAuth, middleware.ParamCheck)
	action.Get("/wol", api.ExecuteWol)
	action.Get("/shutdown", api.ExecuteShutdown)
	action.Get("/ping", api.ExecutePing)

	web := v1.Group("/web")
	web.Post("/login", api.HandleLogin)

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(static.Root),
	}))

	app.Listen(":3000")
}
