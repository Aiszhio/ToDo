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

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedTo time.Time `json:"createdTo"`
}

func getUserInfo(c *fiber.Ctx) (string, error) {
	var req UserNotesRequest
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		fmt.Println("Error unmarshalling request:", err)
		return "", err
	}
	if req.Username == "" {
		return "", errors.New("username is empty")
	}
	if _, err := middleware.GetTokenFromRequest(c); err != nil {
		fmt.Println("Token error:", err)
		return "", err
	}
	return req.Username, nil
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
		if err := conn.QueryRow(c.Context(),
			`SELECT id FROM users WHERE name=$1`, username,
		).Scan(&userId); err != nil {
			fmt.Println("Error getting user ID:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User not found"})
		}

		rows, err := conn.Query(c.Context(),
			`SELECT id, title, created_to
			   FROM notes
			  WHERE user_id=$1
			  ORDER BY created_to
			  LIMIT $2 OFFSET $3`,
			userId, limit, offset,
		)
		if err != nil {
			fmt.Println("Query failed:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database query failed"})
		}
		defer rows.Close()

		var notes []Note
		for rows.Next() {
			var n Note
			if err := rows.Scan(&n.ID, &n.Title, &n.CreatedTo); err != nil {
				fmt.Println("Scan failed:", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read notes"})
			}
			notes = append(notes, n)
		}
		if rows.Err() != nil {
			fmt.Println("Rows iteration error:", rows.Err())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error reading notes"})
		}

		fmt.Printf("Returning notes: %+v\n", notes)

		return c.JSON(fiber.Map{"notes": notes})
	}
}
