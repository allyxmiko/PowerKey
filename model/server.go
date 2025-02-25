package model

type Server struct {
	Port  int    `yaml:"port"`
	Auth  bool   `yaml:"auth"`
	Token string `yaml:"token"`
}
