package main

import (
	"github.com/gofiber/fiber/v2"
)

func regHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func main() {
	webApp := fiber.New()

	webApp.Get("/reg", regHandler)

	webApp.Listen(":5173")
}
