package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email            string            `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password         string            `gorm:"not null"`
	IsAdmin          bool              `gorm:"default:false"`
	UserTestSessions []UserTestSession `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
