package main

import (
	"PowerKey/config"
	"PowerKey/server"
)

func init() {
	config.Init()
}

func main() {
	server.Init()
}
