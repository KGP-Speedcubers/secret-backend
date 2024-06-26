package models

import "gorm.io/gorm"

type Results struct {
	gorm.Model

	UserID   uint      `gorm:"index"` // Foreign key to User.ID
	CompID   uint      `gorm:"index"` // Foreign key to Competition.ID
	Username string    `gorm:"column:username"`
	Event    string    `gorm:"column:event"`
	Ao5      float32   `gorm:"column:ao5"`
	Times    []float32 `gorm:"column:times"`
	Best     float32   `gorm:"column:best"`
}
