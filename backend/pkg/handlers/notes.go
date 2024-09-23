package handlers

import (
	"backend/backend/pkg/middleware"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func ParseToken(tkn string, jwtKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func MakeNote(conn *pgxpool.Pool, secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var storageInf map[string]interface{}
		secretKey := []byte(secret)

		bearerToken, err := middleware.GetTokenFromRequest(c)
		if err != nil {
			return fmt.Errorf("could not get bearer token: %v", err)
		}
		token, err := ParseToken(bearerToken, secretKey)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		var userName string
		fmt.Println(token)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if name, ok := claims["user_name"].(string); ok {
				userName = name
			} else {
				fmt.Println("sosy bebry")
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token: name not found or invalid"})
			}
		}

		err = json.Unmarshal(c.Body(), &storageInf)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return c.SendStatus(fiber.StatusBadRequest)
		}

		note, ok := storageInf["note"].(string)
		if !ok {
			fmt.Println("Note is required:", storageInf["note"])
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Note is required"})
		}

		data, ok := storageInf["date"].(string)
		if !ok {
			fmt.Println("Date is required:", storageInf["date"])
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Date is required"})
		}

		fmt.Println("UserName:", userName, "Note:", note, "Date:", data)

		parseDate, err := time.Parse("02/01/2006", data)
		if err != nil {
			return fmt.Errorf("parse date error: %v", err)
		}
		formatedDate := parseDate.Format("2006-01-02")
		_ = conn.QueryRow(
			context.Background(),
			"INSERT INTO notes(user_id, title, created_to) "+
				"SELECT id, $2, $3 FROM users WHERE name = $1",
			userName, note, formatedDate,
		)

		if err != nil {
			fmt.Println("Failed to insert note:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert note"})
		}
		return c.SendString("You made a note!")

	}
}
