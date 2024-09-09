package handlers

import (
	"backend/backend/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func RegHandler(conn *pgx.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var storage map[string]interface{}

		err := json.Unmarshal(c.Body(), &storage)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(fiber.StatusBadRequest)
		}

		email, emailOk := storage["email"].(string)
		name, nameOk := storage["name"].(string)
		password, passwordOk := storage["password"].(string)

		if !emailOk || !nameOk || !passwordOk {
			fmt.Println("Error here")
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if utils.IsUniqueEmail(email, conn, context.Background()) == true &&
			utils.IsUniqueName(name, conn, context.Background()) == true {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		_, err = conn.Exec(context.Background(), "INSERT INTO users (email, name, password, hash_password) VALUES ($1, $2, $3, $4)", email, name, password, HashPassword(password))
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.SendStatus(fiber.StatusOK)
	}
}
