package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"log"
)

var conn *pgx.Conn

func init() {
	var err error
	conn, err = pgx.Connect(context.Background(), "user=postgres password=5589 host=localhost port=5432 dbname=postgres")
	if err != nil {
		log.Fatal(err)
	}
}

func regHandler(c *fiber.Ctx) error {
	var storage map[string]interface{}

	err := json.Unmarshal(c.Body(), &storage)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	email, emailOk := storage["email"].(string)
	name, nameOk := storage["name"].(string)
	password, passwordOk := storage["password"]

	if !emailOk || !nameOk || !passwordOk {
		fmt.Println("Error here")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO usertodo (email, name, password) VALUES ($1, $2, $3)", email, name, password)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusOK)
}

func main() {
	webApp := fiber.New()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM usertodo")
	if err != nil {
		log.Fatalf("query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		var password string
		var name string
		if err := rows.Scan(&email, &name, &password); err != nil {
			log.Fatal(err)
		}
		fmt.Println(email, name, password)
	}

	webApp.Post("/register", regHandler)
	log.Fatal(webApp.Listen(":5173"))
}
