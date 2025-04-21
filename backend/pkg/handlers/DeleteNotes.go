package handlers

import (
	"backend/backend/pkg/middleware"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DeleteNoteRequest struct {
	ID int `json:"id"`
}

func DeleteNote(conn *pgxpool.Pool, secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bearerToken, err := middleware.GetTokenFromRequest(c)
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

		var req DeleteNoteRequest
		if err = c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		cmd, err := conn.Exec(
			c.Context(),
			`DELETE FROM notes 
             WHERE id = $1 
               AND user_id = (SELECT id FROM users WHERE name = $2)`,
			req.ID, userName,
		)
		if err != nil {
			fmt.Println("Delete failed:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete note"})
		}
		if cmd.RowsAffected() == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found or not yours"})
		}

		return c.SendString("Заметка стёрта!")
	}
}
