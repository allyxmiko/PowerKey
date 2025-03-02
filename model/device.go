package model

type Device struct {
	ID       uint `gorm:"primaryKey"`
	Uid      uint
	IP       string
	Delay    int
	Name     string
	Topic    string
	Password string
	Username string
	Mac      string
}
