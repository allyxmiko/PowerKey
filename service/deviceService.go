package service

import (
	"PowerKey/db"
	"PowerKey/model"
	"PowerKey/model/dto"
	"PowerKey/model/vo"
)

type DeviceService struct {
}

type IDeviceService interface {
	AddDevice(device model.Device) error
	GetDeviceList(uid int) ([]vo.DeviceVo, error)
	GetDeviceById(id int) (vo.DeviceVo, error)
	DeleteDeviceById(id int) error
	UpdateDevice(deviceId int, device dto.DeviceDto) error
}

func NewDeviceService() IDeviceService {
	return &DeviceService{}
}

func (d DeviceService) AddDevice(device model.Device) error {
	return db.DB.Create(&device).Error
}

func (d DeviceService) GetDeviceList(uid int) ([]vo.DeviceVo, error) {
	var devices []vo.DeviceVo
	result := db.DB.Model(&model.Device{}).Where("uid = ?", uid).Find(&devices)
	if result.Error != nil {
		return nil, result.Error
	}
	return devices, nil
}

func (d DeviceService) GetDeviceById(id int) (vo.DeviceVo, error) {
	var device vo.DeviceVo
	result := db.DB.Model(&model.Device{}).Where("id = ?", id).First(&device)
	if result.Error != nil {
		return vo.DeviceVo{}, result.Error
	}
	return device, nil
}

func (d DeviceService) DeleteDeviceById(id int) error {
	return db.DB.Delete(&model.Device{}, id).Error
}
func (d DeviceService) UpdateDevice(deviceId int, device dto.DeviceDto) error {
	return db.DB.Model(&model.Device{ID: deviceId}).Updates(device).Error
}
