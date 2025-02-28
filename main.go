package main

import (
	"PowerKey/config"
	"PowerKey/db"
	"PowerKey/server"
	"log/slog"
)

func init() {
	var err error

	config.Init()

	if err = db.Init(); err != nil {
		slog.Error("初始化数据库失败！", err)
		return
	}
}

func main() {
	server.Init()
}
