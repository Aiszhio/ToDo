package handlers

import (
	"backend/backend/pkg/middleware"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type UpdateNoteRequest struct {
	ID   int    `json:"id"`
	Note string `json:"note"`
	Date string `json:"date"`
}

func UpdateNote(conn *pgxpool.Pool, secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("here!!")
		bearerToken, err := middleware.GetTokenFromRequest(c)
		fmt.Println(bearerToken)
		if err != nil {
			return fmt.Errorf("could not get bearer token: %v", err)
		}
		token, err := ParseToken(bearerToken, []byte(secret))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		var userName string
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if name, ok := claims["user_name"].(string); ok {
				userName = name
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token: name not found"})
			}
		}

		var req UpdateNoteRequest
		if err = c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		t, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format"})
		}
		dateStr := t.Format("2006-01-02")

		cmd, err := conn.Exec(
			c.Context(),
			`UPDATE notes 
             SET title = $1, created_to = $2 
             WHERE id = $3 
               AND user_id = (SELECT id FROM users WHERE name = $4)`,
			req.Note, dateStr, req.ID, userName,
		)
		if err != nil {
			fmt.Println("Update failed:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update note"})
		}
		if cmd.RowsAffected() == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found or not yours"})
		}

		return c.SendString("Заметка обновлена, красавчик!")
	}
}
