package main

import (
	"PowerKey/config"
	"PowerKey/db"
	"PowerKey/server"
	"log/slog"
)

func init() {
	var err error

	err = config.Init()
	if err != nil {
		slog.Error("配置文件初始化失败！", err)
		return
	}

	if err = db.Init(); err != nil {
		slog.Error("初始化数据库失败！", err)
		return
	}
}

func main() {
	server.Init()
}
