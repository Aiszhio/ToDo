package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"log"
)

func EntranceHandler(conn *pgx.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var storage map[string]interface{}
		var requestEmail, requestPassword string

		err := json.Unmarshal(c.Body(), &storage)
		if err != nil {
			return fiber.ErrBadRequest
		}

		requestEmail = storage["email"].(string)
		requestPassword = storage["password"].(string)

		if requestEmail == "" || requestPassword == "" {
			return fiber.ErrBadRequest
		}

		var emailDB, passwordDB string
		var pass bool

		err = conn.QueryRow(context.Background(),
			"SELECT email, hash_password FROM users WHERE email = $1", requestEmail).Scan(&emailDB, &passwordDB)
		if err != nil {
			if err == pgx.ErrNoRows {
				fmt.Println("No rows were returned!")
			} else {
				log.Fatalf("QueryRow failed: %v", err)
			}
		}

		if emailDB == requestEmail {
			pass = true
		}

		if pass {
			if CheckHashPassword(requestPassword, passwordDB) != nil {
				return c.SendStatus(fiber.StatusUnauthorized)
			}
		} else {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.SendStatus(fiber.StatusOK)
	}
}
