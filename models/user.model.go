package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	UserName string `gorm:"username"`
	Email    string `gorm:"email"`
}
