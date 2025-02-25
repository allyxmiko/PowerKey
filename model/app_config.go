package model

type AppConfig struct {
	Server  Server            `yaml:"server"`
	Devices map[string]Device `yaml:"devices"`
}
