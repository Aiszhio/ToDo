package handlers

import (
	"backend/backend/pkg/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type UserNotesRequest struct {
	Username string `json:"userName"`
}

type Notes struct {
	Title     string    `json:"title"`
	CreatedTo time.Time `json:"createdTo"`
}

func getUserInfo(c *fiber.Ctx) (string, error) {
	var RequestName UserNotesRequest

	if err := json.Unmarshal(c.Body(), &RequestName); err != nil {
		fmt.Println("Error in unmarshalling request body")
		return "", err
	}
	if RequestName.Username == "" {
		return "", errors.New("username is empty")
	}

	_, err := middleware.GetTokenFromRequest(c)
	if err != nil {
		fmt.Println("Error in getting token")
		return "", err
	}

	return RequestName.Username, nil
}

func GetUserNotes(conn *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username, err := getUserInfo(c)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		limit := c.QueryInt("limit", 20)
		offset := c.QueryInt("offset", 0)

		var userId int
		err = conn.QueryRow(c.Context(), `SELECT id FROM users WHERE name=$1`, username).Scan(&userId)
		if err != nil {
			fmt.Println("Error in getting user ID:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User not found"})
		}

		rows, err := conn.Query(c.Context(), `SELECT title, created_to FROM notes WHERE user_id=$1 LIMIT $2 OFFSET $3;`, userId, limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database query failed"})
		}
		defer rows.Close()

		var notes []Notes
		for rows.Next() {
			var note Notes
			if err = rows.Scan(&note.Title, &note.CreatedTo); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to scan notes"})
			}
			notes = append(notes, note)
		}
		if rows.Err() != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Rows iteration failed"})
		}

		return c.JSON(fiber.Map{"notes": notes})
	}
}
