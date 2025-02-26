package config

import (
	"PowerKey/model"
	_ "embed"
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path/filepath"
)

//go:embed config.yaml
var defaultConfigContent []byte

var App model.AppConfig
var configPath = "./config"

var validate = validator.New()
var trans, _ = ut.New(zh.New()).GetTranslator("zh")
var _ = zhTrans.RegisterDefaultTranslations(validate, trans)

func Init() error {
	var err error
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err = readConfigFile(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			slog.Info("未找到配置文件...开始创建默认配置文件，请修改后重新启动程序。")
			if err = createConfigFile(); err != nil {
				slog.Error("配置文件生成失败！请自行在程序运行目录创建config目录并且创建config.yaml配置文件。")
			}
		}
	}
	validateConfig()
	return err
}

func readConfigFile() error {
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&App)
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

func validateConfig() {
	count := 0
	slog.Info("开始加载设备列表！")
	for _, device := range App.Devices {
		slog.Info("加载设备配置：" + device.Name)
		if errs := validate.Struct(device); errs != nil {
			for _, err := range errs.(validator.ValidationErrors) {
				count++
				slog.Error(fmt.Sprintf("设备配置：%s，错误字段：%s，原因：%s", device.Name, err.Field(), err.Translate(trans)))
			}
		}
	}
	if count > 0 {
		slog.Error("请按照提示修改错误配置后重新启动程序！")
		os.Exit(1)
	}
	slog.Info("设备列表加载完成！")
}
