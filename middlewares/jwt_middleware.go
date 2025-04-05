package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"os"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")
	if tokenString == "" {
		fmt.Println("no JWT token found, redirecting to login...")
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		fmt.Println("invalid JWT token, redirecting to login...")
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		fmt.Println("invalid JWT claims, redirecting to login...")
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	// Store user ID in locals for access in handlers
	c.Locals("user_id", claims.Issuer)

	return c.Next()
}
