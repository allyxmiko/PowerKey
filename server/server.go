package server

import (
	"PowerKey/server/api"
	"PowerKey/server/static"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/spf13/viper"
	"net/http"
)

func Init() {
	app := fiber.New()

	v1 := app.Group("/api")

	action := v1.Group("/action")
	action.Get("/wol", api.ExecuteWol)
	action.Get("/shutdown", api.ExecuteShutdown)
	action.Get("/ping", api.ExecutePing)

	user := v1.Group("/user")
	user.Post("/login", api.HandleLogin)
	user.Put("/password", api.HandleUpdatePassword)
	user.Put("/token", api.HandleUpdateToken)

	device := v1.Group("/device")
	device.Post("/", api.HandleAddDevice)
	device.Get("/list", api.HandleGetDeviceList)
	//device.Get("/:id", api.HandleGetDevice)
	//device.Put("/:id", api.HandleUpdateDevice)
	//device.Delete("/:id", api.HandleDeleteDevice)

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(static.Root),
	}))

	app.Listen(fmt.Sprintf(":%d", viper.GetInt("server.port")))
}
