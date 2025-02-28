package database

import (
	"PowerKey/global"
	"PowerKey/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func Init() error {
	// TODO 上线后更改为sqlite
	dsn := "power_key_dev:YZDEmtDPYDZw5KAw@tcp(192.168.1.166:3306)/power_key_dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.DB = db

	return nil
}
