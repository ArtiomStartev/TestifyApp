package seed

import (
	"fmt"
	"gorm.io/gorm"
	"testify/models"
)

func SeedAll(db *gorm.DB) {
	seedQuestions(db)
}

func seedQuestions(db *gorm.DB) {
	// Prevent seeding twice
	var count int64
	if db.Model(&models.Question{}).Count(&count); count > 0 {
		return
	}

	questions := []models.Question{
		{
			Text: "Which AWS service is used for object storage?",
			Type: models.SingleChoice,
			Answers: []models.Answer{
				{Text: "Amazon S3", IsCorrect: true},
				{Text: "Amazon EC2", IsCorrect: false},
				{Text: "Amazon RDS", IsCorrect: false},
			},
		},
		{
			Text: "Which of the following are serverless?",
			Type: models.MultipleChoice,
			Answers: []models.Answer{
				{Text: "AWS Lambda", IsCorrect: true},
				{Text: "Amazon EC2", IsCorrect: false},
				{Text: "Amazon DynamoDB", IsCorrect: true},
			},
		},
		{
			Text: "What is the maximum size of an S3 object?",
			Type: models.SingleChoice,
			Answers: []models.Answer{
				{Text: "5 TB", IsCorrect: true},
				{Text: "10 TB", IsCorrect: false},
				{Text: "100 TB", IsCorrect: false},
				{Text: "5 GB", IsCorrect: false},
			},
		},
		{
			Text: "Which AWS service is used for relational databases?",
			Type: models.SingleChoice,
			Answers: []models.Answer{
				{Text: "Amazon RDS", IsCorrect: true},
				{Text: "Amazon DynamoDB", IsCorrect: false},
				{Text: "Amazon S3", IsCorrect: false},
				{Text: "Amazon EC2", IsCorrect: false},
			},
		},
		{
			Text: "Which of the following are NoSQL databases?",
			Type: models.MultipleChoice,
			Answers: []models.Answer{
				{Text: "Amazon DynamoDB", IsCorrect: true},
				{Text: "Amazon RDS", IsCorrect: false},
				{Text: "Amazon Aurora", IsCorrect: false},
				{Text: "Amazon DocumentDB", IsCorrect: true},
			},
		},
	}

	for _, q := range questions {
		if err := db.Create(&q).Error; err != nil {
			fmt.Println("Error seeding question: ", err)
		}
	}
}
