package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"testify/models"
)

type PageController struct {
	DB *gorm.DB
}

func NewPageController(db *gorm.DB) *PageController {
	return &PageController{DB: db}
}

func (pc *PageController) Home(c *fiber.Ctx) error {
	userID := c.Locals("user_id") // Retrieved from middleware

	var user models.User
	if err := pc.DB.First(&user, userID).Error; err != nil {
		fmt.Println("Error fetching user: ", err)
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	return c.Render("templates/home", fiber.Map{
		"Email":   user.Email,
		"IsAdmin": user.IsAdmin,
	})
}

func (pc *PageController) Login(c *fiber.Ctx) error {
	return c.Render("templates/login", fiber.Map{})
}

func (pc *PageController) Signup(c *fiber.Ctx) error {
	return c.Render("templates/signup", fiber.Map{})
}
