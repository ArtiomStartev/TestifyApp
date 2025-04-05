package models

import "gorm.io/gorm"

type UserResponse struct {
	gorm.Model
	UserTestSessionID uint  `gorm:"not null;index"`
	QuestionID        uint  `gorm:"not null;index"`
	AnswerID          *uint `gorm:"index"`
}
