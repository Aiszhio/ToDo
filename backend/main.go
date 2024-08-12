package main

import (
	"backend/backend/pkg/db"
	"backend/backend/pkg/handlers"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	conn, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	webApp := fiber.New()
	webApp.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE",
	}))

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

	webApp.Post("/register", handlers.RegHandler(conn))
	webApp.Post("/entrance", handlers.EntranceHandler(conn))

	log.Fatal(webApp.Listen(":5174"))
}
