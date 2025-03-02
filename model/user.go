package model

type User struct {
	ID       int    `gorm:"autoIncrement"`
	Username string `gorm:"primaryKey"`
	Token    string
	Password string
}
