package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAuthCookie(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}

		key := os.Getenv("SECRET")
		return []byte(key), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"err":     true,
			"message": "You cannot access this resource",
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expiryTime := claims["exp"].(float64)

		if time.Now().Unix() > int64(expiryTime) {
			c.Cookie(&fiber.Cookie{
				Name:  "token",
				Value: "",
			})

			return c.SendStatus(fiber.StatusUnauthorized)
		}
	} else {
		log.Println("error when parsing claims")
	}

	return c.Next()
}
