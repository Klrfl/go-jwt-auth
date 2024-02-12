package main

import (
	"klrfl/go-jwt-auth/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Init()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello!",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"err":     false,
			"message": "You are in login endpoint",
		})
	})

	app.Listen(":8080")
}
