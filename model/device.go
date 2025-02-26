package model

type Device struct {
	Name     string `yaml:"name" validate:"required"`
	IP       string `yaml:"ip" validate:"required,ip"`
	Mac      string `yaml:"mac" validate:"required,mac"`
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Delay    int    `yaml:"delay" validate:"required,gte=0"`
}
