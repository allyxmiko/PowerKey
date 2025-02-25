package main

import (
	"PowerKey/config"
	"PowerKey/devices"
	"PowerKey/server"
	"log/slog"
)

func init() {
	var err error
	if err = config.Init(); err != nil {
		slog.Error("初始化配置文件失败！", err)
	}
	devices.Init()
}

func main() {
	server.Init()
}
