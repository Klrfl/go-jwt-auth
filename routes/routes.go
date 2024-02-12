package routes

import (
	"klrfl/go-jwt-auth/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(app *fiber.App) {
	app.Use(logger.New())
	base := app.Group("/api")

	base.Get("/", handlers.Base)
	base.Post("/login", handlers.Login)
	base.Post("/sign-in", handlers.Signin)
	base.Get("/public", handlers.Public)
	base.Get("/private", handlers.Private)
}
