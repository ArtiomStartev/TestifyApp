package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"testify/controllers"
	"testify/middlewares"
)

func Setup(app *fiber.App, db *gorm.DB) {
	pc := controllers.NewPageController(db)
	app.Get("/", middlewares.AuthMiddleware, pc.Home)
	app.Get("/signup", pc.Signup)
	app.Get("/login", pc.Login)

	ac := controllers.NewAuthController(db)
	app.Post("/signup", ac.Signup)
	app.Post("/login", ac.Login)
	app.Get("/logout", ac.Logout)

	tc := controllers.NewTestController(db)
	app.Get("/test", middlewares.AuthMiddleware, tc.StartTest)
	app.Post("/submit-test", middlewares.AuthMiddleware, tc.SubmitTest)
	app.Get("/test-result/:session_id", middlewares.AuthMiddleware, tc.ShowResult)
	app.Get("/results", middlewares.AuthMiddleware, tc.UserResults)
	app.Get("/all-users-results", middlewares.AuthMiddleware, middlewares.AdminMiddleware(db), tc.AllUsersResults)
}
