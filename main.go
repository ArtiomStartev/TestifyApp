package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"net/http"
	"testify/database"
	"testify/helpers"
	"testify/routes"
)

//go:embed templates/*
var templatesFS embed.FS

var (
	db *gorm.DB
)

func main() {
	engine := html.NewFileSystem(http.FS(templatesFS), ".gohtml")
	helpers.RegisterCustomTemplateFuncs(engine)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.Setup(app, db)

	if err := app.Listen(":80"); err != nil {
		panic("Error starting server on port 80: " + err.Error())
	}
}

func init() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	db, err = database.NewDBConn()
	if err != nil {
		panic("Error connecting to database: " + err.Error())
	}
}
