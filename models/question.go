package models

import (
	"gorm.io/gorm"
)

type QuestionType string

const (
	SingleChoice   QuestionType = "single"
	MultipleChoice QuestionType = "multiple"
)

type Question struct {
	gorm.Model
	Text    string       `gorm:"not null"`
	Type    QuestionType `gorm:"type:enum('single','multiple');not null"`
	Answers []Answer     `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE"`
}
