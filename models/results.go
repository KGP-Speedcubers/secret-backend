package models

import "gorm.io/gorm"

type Results struct {
	gorm.Model

	UserID   uint   `gorm:"index"` // Foreign key to User.ID
	CompID   uint   `gorm:"index"` // Foreign key to Competition.ID
	Username string `gorm:"column:username"`
	Event    string `gorm:"column:event"`
	Ao5      string `gorm:"column:ao5"`
	Times    string `gorm:"column:times"`
	Best     string `gorm:"column:best"`
}
