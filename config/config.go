package config

import (
	"PowerKey/utils"
	"errors"
	"github.com/spf13/viper"
	"os"
)

var Username = "admin"
var configPath = "./config/config.yaml"

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		viper.SetDefault("server.port", 3000)
		viper.SetDefault("jwt.secret", utils.RandomString(32))
		viper.SetDefault("jwt.expire", 168)
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				file, err := os.Create(configPath)
				if err != nil {
					return err
				}
				defer file.Close()
			}
			err = viper.WriteConfig()
		}
		return err
	}

	return nil
}
