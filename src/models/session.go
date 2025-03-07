package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserID  uint      `gorm:"not null"`
	Token   string    `gorm:"unique;not null"`
	Expired time.Time `gorm:"not null"`
}
