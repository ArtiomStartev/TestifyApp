package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"strconv"
	"testify/helpers"
	"testify/models"
	"time"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ac *AuthController) Signup(c *fiber.Ctx) error {
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.HandleError(c, fiber.StatusBadRequest, err, "Invalid request")
	}

	if req.Email == "" || req.Password == "" {
		return helpers.HandleError(c, fiber.StatusBadRequest, nil, "Email and password are required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, err, "Error hashing password")
	}

	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err = ac.DB.Create(&user).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, err, "Error creating user")
	}

	return c.Redirect("/login", fiber.StatusSeeOther)
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.HandleError(c, fiber.StatusBadRequest, err, "Invalid request")
	}

	if req.Email == "" || req.Password == "" {
		return helpers.HandleError(c, fiber.StatusBadRequest, nil, "Email and password are required")
	}

	var user models.User
	if err := ac.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return helpers.HandleError(c, fiber.StatusUnauthorized, err, "Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return helpers.HandleError(c, fiber.StatusUnauthorized, err, "Invalid credentials")
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		return helpers.HandleError(c, fiber.StatusInternalServerError, err, "Error generating JWT token")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true, // Set to true in production
		SameSite: "Strict",
	})

	return c.Redirect("/", fiber.StatusSeeOther)
}

func (ac *AuthController) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   true, // Set to true in production
		SameSite: "Strict",
	})

	return c.Redirect("/login", fiber.StatusSeeOther)
}

// Generate JWT Token
func generateJWT(userID uint) (string, error) {
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userID)),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
