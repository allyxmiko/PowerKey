package api

import (
	"PowerKey/config"
	"PowerKey/model"
	"PowerKey/model/dto"
	"PowerKey/server/resp"
	"PowerKey/validator"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HandleAddDevice(c *fiber.Ctx) error {
	var deviceDto dto.DeviceDto
	var user model.User
	var err error

	if err = c.BodyParser(&deviceDto); err != nil {
		return c.JSON(resp.Err(resp.InvalidQueryParam))
	}
	if errs := validator.Validate.Struct(&deviceDto); errs != nil {
		for _, err = range errs.(validate.ValidationErrors) {
			return c.JSON(resp.Data(err))
		}
	}
	if user, err = userService.FindUserByUsername(deviceDto.Username); err != nil {
		return c.JSON(resp.Err(resp.UserNotFound))
	}
	device := model.Device{
		Delay:    deviceDto.Delay,
		IP:       deviceDto.IP,
		Mac:      deviceDto.Mac,
		Name:     deviceDto.Name,
		Password: deviceDto.Password,
		Topic:    deviceDto.Topic,
		Username: deviceDto.Username,
		Uid:      user.ID,
	}
	if err = deviceService.AddDevice(device); err != nil {
		return c.JSON(resp.Err(resp.DeviceAddFailed))
	}
	return c.JSON(resp.Ok("设备添加成功！"))
}

func HandleGetDeviceList(c *fiber.Ctx) error {
	var user model.User
	var result []model.Device
	var err error
	if user, err = userService.FindUserByUsername(config.Username); err != nil {
		return c.JSON(resp.Err(resp.UserNotFound))
	}
	result, err = deviceService.GetDeviceList(user.ID)
	if err != nil {
		return c.JSON(resp.Err(resp.DeviceNotFound))
	}
	return c.JSON(resp.Data(result))
}

//func HandleGetDevice(ctx *fiber.Ctx) error {
//
//}
//
//func HandleUpdateDevice(ctx *fiber.Ctx) error {
//
//}
//
//func HandleDeleteDevice(ctx *fiber.Ctx) error {
//
//}
