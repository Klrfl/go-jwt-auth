package handlers

import (
	"klrfl/go-jwt-auth/database"
	"klrfl/go-jwt-auth/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Base(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "hello!",
	})
}

func Login(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var incomingUser models.User

	if err := c.BodyParser(&incomingUser); err != nil {
		return c.JSON(fiber.Map{
			"err":     true,
			"message": "something wrong with payload",
		})
	}

	var existingUser models.User
	result := database.DB.
		Limit(1).
		Where("email = ?", incomingUser.Email).
		Find(&existingUser)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err":     true,
			"message": "error when querying database",
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"err":     false,
			"message": "user doesn't exist",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(incomingUser.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err":     true,
			"message": "wrong email or password",
		})
	}

	//TODO: make JWT token
	expirationTime := time.Now().Add(1 * time.Hour).UTC()
	now := time.Now().UTC()
	key := os.Getenv("SECRET")
	claims := &models.JWTClaim{
		Name:  incomingUser.Name,
		Email: incomingUser.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   incomingUser.ID.String(),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err":     true,
			"message": "error when authenticating",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   int(expirationTime.Unix()),
		HTTPOnly: false,
		Secure:   true,
	})

	return c.JSON(fiber.Map{
		"err":     false,
		"message": "succesfully logged in",
	})
}

func Signup(c *fiber.Ctx) error {
	c.Accepts("applicaton/json")

	var newUser models.User

	if err := c.BodyParser(&newUser); err != nil {
		return c.JSON(fiber.Map{
			"err":     true,
			"message": "something wrong with payload",
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)

	if err != nil {
		return c.JSON(fiber.Map{
			"err":     true,
			"message": "error when generating new user",
		})
	}

	newUser.Password = string(password)
	database.DB.Create(&newUser)

	return c.JSON(fiber.Map{
		"err":     false,
		"message": "sign up successful",
	})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
	})

	return c.JSON(fiber.Map{
		"err":     false,
		"message": "logged out successfully",
	})
}

func Public(c *fiber.Ctx) error {
	return c.SendString("Public string")
}

func Private(c *fiber.Ctx) error {
	return c.SendString("Private string. You can only see this if you're authenticated")
}
