package utils

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
	"time"
)

type Notification struct {
	Email    string
	Password string
}

func MessageLimit(limiter *rate.Limiter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !limiter.AllowN(time.Now(), 1) {
			return c.SendStatus(fiber.StatusTooManyRequests)
		}
		return c.Next()
	}
}
