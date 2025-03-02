package service

import (
	"PowerKey/db"
	"PowerKey/model"
)

type DeviceService struct {
}

type IDeviceService interface {
	AddDevice(device model.Device) error
	GetDeviceList(uid int) ([]model.Device, error)
	GetDeviceById(id int) (model.Device, error)
	DeleteDeviceById(id int) error
	UpdateDevice(deviceId int, device model.Device) error
}

func NewDeviceService() IDeviceService {
	return &DeviceService{}
}

func (d DeviceService) AddDevice(device model.Device) error {
	return db.DB.Create(&device).Error
}

func (d DeviceService) GetDeviceList(uid int) ([]model.Device, error) {
	var devices []model.Device
	result := db.DB.Model(&model.Device{}).Where("uid = ?", uid).Find(&devices)
	if result.Error != nil {
		return nil, result.Error
	}
	return devices, nil
}

func (d DeviceService) GetDeviceById(id int) (model.Device, error) {
	var device model.Device
	result := db.DB.Model(&model.Device{}).Where("id = ?", id).First(&device)
	if result.Error != nil {
		return model.Device{}, result.Error
	}
	return device, nil
}

func (d DeviceService) DeleteDeviceById(id int) error {
	return db.DB.Delete(&model.Device{}, id).Error
}
func (d DeviceService) UpdateDevice(deviceId int, device model.Device) error {
	return db.DB.Model(&model.Device{ID: deviceId}).Updates(device).Error
}
