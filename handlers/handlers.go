package handlers

import "github.com/gofiber/fiber/v2"

func Base(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "hello!",
	})
}

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{

		"err":     false,
		"message": "You are in login endpoint",
	})
}

func Signin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{

		"err":     false,
		"message": "You are in signin endpoint",
	})
}

func Public(c *fiber.Ctx) error {
	return c.SendString("Public string")
}

func Private(c *fiber.Ctx) error {
	return c.SendString("Private string. You can only see this if you're authenticated")
}
