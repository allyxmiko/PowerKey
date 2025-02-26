package model

type Server struct {
	Port  int    `yaml:"port" validate:"required,hostname_port"`
	Auth  bool   `yaml:"auth" validate:"required"`
	Token string `yaml:"token"`
}
