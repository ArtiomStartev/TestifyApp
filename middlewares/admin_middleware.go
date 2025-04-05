package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"testify/helpers"
	"testify/models"
)

func AdminMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id") // Retrieved from middleware

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			return helpers.HandleError(c, fiber.StatusNotFound, err, "User not found")
		}

		if !user.IsAdmin {
			return helpers.HandleError(c, fiber.StatusForbidden, nil, "Access denied")
		}

		return c.Next()
	}
}
