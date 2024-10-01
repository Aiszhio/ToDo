package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

func GetTokenFromRequest(c *fiber.Ctx) (string, error) {
	header := c.Get("Authorization")
	if header == "" {
		return "", errors.New("Authorization header is empty")
	}

	headersPart := strings.Split(header, " ")
	if len(headersPart) != 2 || headersPart[0] != "Bearer" {
		return "", errors.New("Authorization header is not Bearer")
	}

	if len(headersPart[1]) == 0 {
		return "", errors.New("Token is empty")
	}
	return headersPart[1], nil
}

func SaveToken(ctx context.Context, rdb *redis.Client, userID, token string) error {
	err := rdb.Set(ctx, userID, token, 10*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("Error saving token: %v", err)
	}
	return nil
}

func GetToken(c *fiber.Ctx, rdb *redis.Client, userID string) string {
	token, err := rdb.Get(c.Context(), userID).Result()
	if err != nil {
		return ""
	} else if token == "" {
		return ""
	}
	return token
}
