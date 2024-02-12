package routes

import (
	"klrfl/go-jwt-auth/handlers"
	"klrfl/go-jwt-auth/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(app *fiber.App) {
	app.Use(limiter.New())
	app.Use(logger.New())
	base := app.Group("/api")

	base.Get("/", handlers.Base)
	base.Post("/login", handlers.Login)
	base.Post("/sign-up", handlers.Signup)
	base.Post("/logout", handlers.Logout)
	base.Get("/public", handlers.Public)

	base.Use(middleware.CheckAuthCookie)
	base.Get("/private", handlers.Private)
}
