package model

type Device struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Uid      int    `json:"-"`
	IP       string `json:"ip" validate:"required,ip"`
	Delay    int    `json:"delay" validate:"required,gte=0"`
	Name     string `json:"name" validate:"required"`
	Topic    string `json:"topic" gorm:"unique" validate:"required,alphanum"`
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
	Mac      string `json:"mac"  validate:"required,mac"`
}
