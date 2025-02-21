package main

import (
	"PowerKey/config"
	"fmt"
)

func init() {
	err := config.Init()
	fmt.Println(err)
}

func main() {
	println(config.App.Server.Token)
}
