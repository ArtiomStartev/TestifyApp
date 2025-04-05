package models

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	QuestionID uint   `gorm:"not null;index"`
	Text       string `gorm:"not null"`
	IsCorrect  bool   `gorm:"default:false"`
}
