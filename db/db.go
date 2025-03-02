package db

import (
	"PowerKey/model"
	"PowerKey/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"time"
)

var DB *gorm.DB

func Init() error {
	dbPath := "./data/db/PowerKey.db"
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return err
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.User{}, &model.Device{})
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

	DB = db

	if err := db.First(&model.User{Username: "admin"}); err != nil {
		db.Create(&model.User{
			Username: "admin",
			Token:    utils.RandomString(32),
			Password: utils.HashPassword("admin"),
		})
	}

	return nil
}
