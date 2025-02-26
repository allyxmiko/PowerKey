package model

type AppConfig struct {
	Server  Server   `yaml:"server"`
	Devices []Device `yaml:"devices"`
}
