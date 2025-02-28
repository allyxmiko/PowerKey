package model

type User struct {
	ID       uint   `gorm:"autoIncrement"`
	Username string `gorm:"primaryKey"`
	Token    string
	Password string
}
