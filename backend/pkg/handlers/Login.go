package handlers

import (
	"backend/backend/pkg/middleware"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type User struct {
	Id   string
	Name string
}

var WebUser User

var storageLog map[string]interface{}

var Token string

func EntranceHandler(conn *pgxpool.Pool, rdb *redis.Client, secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestEmail, requestPassword string

		err := json.Unmarshal(c.Body(), &storageLog)
		if err != nil {
			return fiber.ErrBadRequest
		}

		requestEmail = storageLog["email"].(string)
		requestPassword = storageLog["password"].(string)

		if requestEmail == "" || requestPassword == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		var idDB, emailDB, nameDB, passwordDB string

		err = conn.QueryRow(context.Background(),
			"SELECT id, email, name, hash_password FROM users WHERE email = $1", requestEmail).Scan(&idDB,
			&emailDB, &nameDB, &passwordDB)
		if err != nil {
			if err == pgx.ErrNoRows {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "WebUser not found"})
			} else {
				log.Printf("QueryRow failed: %v", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
			}
		}

		WebUser.Id = idDB
		WebUser.Name = nameDB

		if emailDB == requestEmail && CheckHashPassword(requestPassword, passwordDB) == nil {
			secretKey := []byte(secret)
			Token, err = JWTGenerator(secretKey)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Error while signing Token")
			}
			err = middleware.SaveToken(context.Background(), rdb, WebUser.Id, Token)
			if err != nil {
				log.Fatalf("failed to save token: %v", err)
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"Token": Token})
		} else {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}
}

func JWTGenerator(secret []byte) (string, error) {
	if WebUser.Id == "" || WebUser.Name == "" {
		return "", fmt.Errorf("invalid WebUser data")
	}

	claims := jwt.MapClaims{
		"user_id":   WebUser.Id,
		"user_name": WebUser.Name,
		"exp":       time.Now().Add(time.Minute * 10).Unix(),
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := Token.SignedString(secret)
	if err != nil {
		return "", err
	}
	fmt.Println(Token)
	return signedToken, nil
}
