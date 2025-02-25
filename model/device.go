package model

type Device struct {
	IP       string `yaml:"ip"`
	Mac      string `yaml:"mac"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Delay    int    `yaml:"delay"`
}
