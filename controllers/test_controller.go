package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
	"slices"
	"strconv"
	"testify/helpers"
	"testify/models"
)

type SubmitTestRequest struct {
	UserID           uint                `json:"user_id"`
	Answers          map[string][]string `json:"answers"` // map[questionID][]answerIDs
	TimeTakenSeconds int                 `json:"time_taken_seconds"`
}

type UserResults struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	Sessions []models.UserTestSession
}

type TestController struct {
	DB *gorm.DB
}

func NewTestController(db *gorm.DB) *TestController {
	return &TestController{DB: db}
}

func (tc *TestController) StartTest(c *fiber.Ctx) error {
	userID := c.Locals("user_id") // Retrieved from middleware

	var user models.User
	var questions []models.Question

	if err := tc.DB.First(&user, userID).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusNotFound, err, "User not found")
	}

	if err := tc.DB.Preload("Answers").Find(&questions).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, err, "Failed to fetch questions")
	}

	return c.Render("templates/test", fiber.Map{
		"UserID":    user.ID,
		"IsAdmin":   user.IsAdmin,
		"Questions": questions,
	})
}

func (tc *TestController) SubmitTest(c *fiber.Ctx) error {
	var req SubmitTestRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.HandleError(c, fiber.StatusBadRequest, err, "Invalid request")
	}

	session := models.UserTestSession{
		UserID:          req.UserID,
		DurationSeconds: req.TimeTakenSeconds,
	}
	if err := tc.DB.Create(&session).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, err, "Failed to create test session")
	}

	var totalCorrect int
	for questionIDStr, answerIDs := range req.Answers {
		questionID, err := strconv.Atoi(questionIDStr)
		if err != nil {
			fmt.Println("Error converting question ID: ", err)
			continue
		}

		var correctAnswerIDs []uint
		err = tc.DB.Model(&models.Answer{}).
			Select("id").
			Where("question_id = ? AND is_correct = true", questionID).
			Pluck("id", &correctAnswerIDs).Error

		if err != nil {
			fmt.Println("Error fetching correct answers: ", err)
			continue
		}

		isCorrect := len(correctAnswerIDs) == len(answerIDs)
		for _, answerIDStr := range answerIDs {
			answerID, err := strconv.Atoi(answerIDStr)
			if err != nil {
				fmt.Println("Error converting answer ID: ", err)
				continue
			}
			answerIDUint := uint(answerID)

			if isCorrect && !slices.Contains(correctAnswerIDs, answerIDUint) {
				isCorrect = false
			}

			response := models.UserResponse{
				UserTestSessionID: session.ID,
				QuestionID:        uint(questionID),
				AnswerID:          &answerIDUint,
			}
			if err = tc.DB.Create(&response).Error; err != nil {
				fmt.Println("Error creating user response: ", err)
				continue
			}
		}

		if isCorrect {
			totalCorrect++
		}
	}

	var totalQuestions int64
	if err := tc.DB.Model(&models.Question{}).Count(&totalQuestions).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, err, "Failed to count total questions")
	}

	session.Score = totalCorrect
	session.Passed = totalCorrect >= int(math.Ceil(float64(totalQuestions)/2))
	if err := tc.DB.Save(&session).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, err, "Failed to update test session score")
	}

	return c.Redirect(fmt.Sprintf("/test-result/%d", session.ID), fiber.StatusSeeOther)
}

func (tc *TestController) ShowResult(c *fiber.Ctx) error {
	userID := c.Locals("user_id") // Retrieved from middleware
	sessionID := c.Params("session_id")

	var user models.User
	if err := tc.DB.First(&user, userID).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusNotFound, err, "User not found")
	}

	var session models.UserTestSession
	if err := tc.DB.Where("id = ? AND user_id = ?", sessionID, userID).First(&session).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusNotFound, err, "Test session not found")
	}

	var totalQuestions int64
	if err := tc.DB.Model(&models.Question{}).Count(&totalQuestions).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, err, "Failed to count total questions")
	}

	return c.Render("templates/result", fiber.Map{
		"IsAdmin":        user.IsAdmin,
		"TotalQuestions": totalQuestions,
		"TotalCorrect":   session.Score,
		"Passed":         session.Passed,
		"TimeTaken":      session.DurationSeconds,
	})
}

func (tc *TestController) UserResults(c *fiber.Ctx) error {
	userID := c.Locals("user_id") // Retrieved from middleware

	var user models.User
	if err := tc.DB.First(&user, userID).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusNotFound, err, "User not found")
	}

	var sessions []models.UserTestSession
	if err := tc.DB.Where("user_id = ?", user.ID).Find(&sessions).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, err, "Failed to fetch test sessions")
	}

	return c.Render("templates/user_results", fiber.Map{
		"IsAdmin":  user.IsAdmin,
		"Sessions": sessions,
	})
}

func (tc *TestController) AllUsersResults(c *fiber.Ctx) error {
	var users []models.User
	if err := tc.DB.Find(&users).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, err, "Failed to fetch users")
	}

	var usersResults []UserResults
	for _, user := range users {
		var sessions []models.UserTestSession
		if err := tc.DB.Where("user_id = ?", user.ID).Find(&sessions).Error; err != nil {
			return helpers.HandleError(c, fiber.StatusInternalServerError, err, "Failed to fetch test sessions")
		}
		usersResults = append(usersResults, UserResults{
			UserID:   user.ID,
			Email:    user.Email,
			Sessions: sessions,
		})
	}

	return c.Render("templates/all_users_results", fiber.Map{
		"IsAdmin":      true,
		"UsersResults": usersResults,
	})
}
