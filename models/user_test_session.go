package models

import (
	"gorm.io/gorm"
)

type UserTestSession struct {
	gorm.Model
	UserID          uint           `gorm:"not null;index"`
	DurationSeconds int            `gorm:"not null"`
	Score           int            `gorm:"not null"`
	Passed          bool           `gorm:"not null;default:false"`
	Responses       []UserResponse `gorm:"foreignKey:UserTestSessionID;constraint:OnDelete:CASCADE"`
}
