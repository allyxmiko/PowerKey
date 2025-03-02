package api

import (
	"PowerKey/config"
	"PowerKey/model"
	"PowerKey/server/resp"
	"PowerKey/validator"
	"github.com/gofiber/fiber/v2"
)

func HandleAddDevice(c *fiber.Ctx) error {
	var device model.Device

	if err := c.BodyParser(&device); err != nil {
		return c.JSON(resp.Err(resp.InvalidQueryParam))
	}

	if errsMap, err := validator.Struct(&device); err != nil {
		return c.JSON(resp.Data(errsMap))
	}

	if user, err := userService.FindUserByUsername(config.Username); err != nil {
		return c.JSON(resp.Err(resp.UserNotFound))
	} else {
		device.ID = user.ID
	}

	if err := deviceService.AddDevice(device); err != nil {
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

func HandleGetDevice(c *fiber.Ctx) error {
	deviceId, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(resp.Err(resp.InvalidPathParam))
	}
	device, err := deviceService.GetDeviceById(deviceId)
	if err != nil {
		return c.JSON(resp.Err(resp.DeviceNotFound))
	}
	return c.JSON(resp.Data(device))
}

func HandleUpdateDevice(c *fiber.Ctx) error {
	var device model.Device
	deviceId, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(resp.Err(resp.InvalidPathParam))
	}
	if device, err = deviceService.GetDeviceById(deviceId); err != nil {
		return c.JSON(resp.Err(resp.DeviceNotFound))
	}
	if err = c.BodyParser(&device); err != nil {
		return c.JSON(resp.Err(resp.InvalidQueryParam))
	}

	if errsMap, err := validator.Struct(&device); err != nil {
		return c.JSON(resp.Data(errsMap))
	}
	if err := deviceService.UpdateDevice(deviceId, device); err != nil {
		return c.JSON(resp.Err(resp.DeviceUpdateFailed))
	}
	return c.JSON(resp.Ok("设备更新成功！"))
}

func HandleDeleteDevice(c *fiber.Ctx) error {
	deviceId, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(resp.Err(resp.InvalidPathParam))
	}
	_, err = deviceService.GetDeviceById(deviceId)
	if err != nil {
		return c.JSON(resp.Err(resp.DeviceNotFound))
	}
	if err = deviceService.DeleteDeviceById(deviceId); err != nil {
		return c.JSON(resp.Err(resp.DeviceDeleteFailed))
	}
	return c.JSON(resp.Ok("设备删除成功！"))
}
