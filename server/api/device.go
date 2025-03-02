package api

import (
	"PowerKey/config"
	"PowerKey/model"
	"PowerKey/model/dto"
	"PowerKey/model/vo"
	"PowerKey/server/resp"
	"PowerKey/validator"
	"github.com/gofiber/fiber/v2"
)

func HandleAddDevice(c *fiber.Ctx) error {
	var deviceDto dto.DeviceDto
	var user model.User
	var err error

	if err = c.BodyParser(&deviceDto); err != nil {
		return c.JSON(resp.Err(resp.InvalidQueryParam))
	}

	if errsMap, err := validator.Struct(&deviceDto); err != nil {
		return c.JSON(resp.Data(errsMap))
	}

	if user, err = userService.FindUserByUsername(config.Username); err != nil {
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
	var result []vo.DeviceVo
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
	var deviceDto dto.DeviceDto
	deviceId, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(resp.Err(resp.InvalidPathParam))
	}
	_, err = deviceService.GetDeviceById(deviceId)
	if err != nil {
		return c.JSON(resp.Err(resp.DeviceNotFound))
	}
	if err = c.BodyParser(&deviceDto); err != nil {
		return c.JSON(resp.Err(resp.InvalidQueryParam))
	}

	if errsMap, err := validator.Struct(&deviceDto); err != nil {
		return c.JSON(resp.Data(errsMap))
	}
	if err := deviceService.UpdateDevice(deviceId, deviceDto); err != nil {
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
