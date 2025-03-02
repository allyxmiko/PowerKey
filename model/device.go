package model

type Device struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Uid      int    `json:"-" json:"uid"`
	IP       string `json:"ip" validate:"ip"`
	Delay    int    `json:"delay" validate:"gte=0"`
	Name     string `json:"name"`
	Topic    string `json:"topic" validate:"alphanum"`
	Password string `json:"password"`
	Username string `json:"username"`
	Mac      string `json:"mac"  validate:"mac"`
}
