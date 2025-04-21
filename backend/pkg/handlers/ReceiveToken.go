package handlers

import (
	"backend/backend/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ReceiveToken(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := middleware.GetToken(c, rdb, WebUser.Id)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
	}
}
