package api

import (
	"PowerKey/action"
	"PowerKey/action/wol"
	"PowerKey/model"
	"PowerKey/server/resp"
	"github.com/gofiber/fiber/v2"
)

func ExecuteWol(c *fiber.Ctx) error {
	device := c.Locals("device").(model.Device)
	if device.Mac == "" {
		return c.JSON(resp.Err(resp.MacNotFound))
	}
	if err := wol.WakeOnLAN(device.Mac); err != nil {
		return c.JSON(resp.Err(resp.ActionExecuteFailed))
	}
	return c.JSON(resp.Ok("执行成功！"))
}

func ExecuteShutdown(c *fiber.Ctx) error {
	device := c.Locals("device").(model.Device)
	if device.IP == "" || device.Username == "" || device.Password == "" {
		return c.JSON(resp.Err(resp.MissingConfiguration))
	}
	if err := action.Shutdown(device.Username, device.Password, device.IP, device.Delay); err != nil {
		return c.JSON(resp.Err(resp.ActionExecuteFailed))
	}
	return c.JSON(resp.Ok("执行成功！"))
}

func ExecutePing(c *fiber.Ctx) error {
	device := c.Locals("device").(model.Device)
	if device.IP == "" {
		return c.JSON(resp.Err(resp.MissingConfiguration))
	}
	if err := action.Ping(device.IP); err != nil {
		return c.JSON(resp.Err(resp.DeviceOffline))
	}
	return c.JSON(resp.Ok("当前设备在线！"))
}
