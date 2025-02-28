package main

import (
	"PowerKey/config"
	"PowerKey/database"
	"PowerKey/server"
	"log/slog"
)

func init() {
	var err error
	if err = config.Init(); err != nil {
		slog.Error("初始化配置文件失败！", err)
		return
	}
	err = database.Init()
	if err != nil {
		slog.Error("初始化数据库失败！", err)
		return
	}
}

func main() {
	server.Init()
}
