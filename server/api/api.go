package api

import "PowerKey/service"

var (
	userService   = service.NewUserService()
	deviceService = service.NewDeviceService()
)
