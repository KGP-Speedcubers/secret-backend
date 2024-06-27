package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Email    string `gorm:"column:email"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}
