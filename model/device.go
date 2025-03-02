package model

type Device struct {
	ID       int `gorm:"primaryKey"`
	Uid      int
	IP       string
	Delay    int
	Name     string
	Topic    string
	Password string
	Username string
	Mac      string
}
