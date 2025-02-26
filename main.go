package main

import (
	"PowerKey/config"
	"PowerKey/server"
	"log/slog"
	"os"
)

func init() {
	var err error
	if err = config.Init(); err != nil {
		slog.Error("初始化配置文件失败！", err)
		os.Exit(1)
	}
}

func main() {
	server.Init()
}
