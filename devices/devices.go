package devices

import (
	"PowerKey/config"
	"log/slog"
)

func Init() {
	slog.Info("开始加载设备列表！")
	for k := range config.App.Devices {
		slog.Info("加载设备：" + k)
	}
}
