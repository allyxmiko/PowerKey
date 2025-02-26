package server

import (
	"PowerKey/server/api"
	"PowerKey/server/middleware"
	"PowerKey/server/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
)

func Init() {
	app := fiber.New()

	v1 := app.Group("/api", middleware.BaseAuth, middleware.ParamCheck)
	action := v1.Group("/action")
	action.Get("/wol", api.ExecuteWol)
	action.Get("/shutdown", api.ExecuteShutdown)
	action.Get("/ping", api.ExecutePing)

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(web.Root),
	}))

	app.Listen(":3000")
}
