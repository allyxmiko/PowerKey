package service

import (
	"PowerKey/db"
	"PowerKey/model"
)

type DeviceService struct {
}

type IDeviceService interface {
	AddDevice(device model.Device) error
	GetDeviceList(uid uint) ([]model.Device, error)
}

func NewDeviceService() IDeviceService {
	return &DeviceService{}
}

func (d DeviceService) AddDevice(device model.Device) error {
	return db.DB.Create(&device).Error
}

func (d DeviceService) GetDeviceList(uid uint) ([]model.Device, error) {
	var devices []model.Device
	result := db.DB.Where("uid = ?", uid).Find(&devices)
	if result.Error != nil {
		return nil, result.Error
	}
	return devices, nil
}
