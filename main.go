package main

import (
	"klrfl/go-jwt-auth/database"
	"klrfl/go-jwt-auth/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Init()

	app := fiber.New()
	routes.Init(app)
	app.Listen(":8080")
}
