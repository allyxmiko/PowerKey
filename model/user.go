package model

type User struct {
	ID       int    `gorm:"autoIncrement" json:"id,omitempty"`
	Username string `gorm:"primaryKey" json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
	Password string `json:"password,omitempty"`
}
