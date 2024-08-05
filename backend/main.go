package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func regHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func main() {
	webApp := fiber.New()

	webApp.Get("/registration", regHandler)
	http.ListenAndServe(":8080", nil)
}
