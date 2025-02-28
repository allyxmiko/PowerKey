package config

import "flag"

var Port int

func Init() {
	flag.IntVar(&Port, "port", 3000, "服务端口（默认 3000）")

	flag.Parse()
}
