package config

import (
	"PowerKey/model"
	_ "embed"
	"errors"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path/filepath"
)

//go:embed config.yaml
var defaultConfigContent []byte

var App model.AppConfig
var configPath = getExecutablePath()

func Init() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			slog.Info("未找到配置文件...已生成默认配置文件，请修改后重新启动程序。")
			if err := createConfigFile(); err != nil {
				return err
			}
		}
		return err
	}
	if err := viper.Unmarshal(&App); err != nil {
		return err
	}
	return nil
}

func createConfigFile() error {
	file, err := os.Create(filepath.Join(configPath, "config.yaml"))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(defaultConfigContent)
	if err != nil {
		return err
	}
	return nil
}

func getExecutablePath() string {
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	return exeDir
}
