package helpers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"html/template"
)

// RegisterCustomTemplateFuncs registers custom template functions
func RegisterCustomTemplateFuncs(engine *html.Engine) {
	// Add custom functions here
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		// Add more functions as needed
		// Example: "multiply": func(a, b int) int { return a * b }
	}

	// Register the functions with the engine
	engine.AddFuncMap(funcMap)
}

// HandleError is a utility function to handle errors and send a response
func HandleError(c *fiber.Ctx, statusCode int, err error, message string) error {
	fmt.Printf("%s: %v", message, err)
	return c.Status(statusCode).SendString(message)
}
