package main

import (
	"backend/backend/pkg/db"
	"backend/backend/pkg/handlers"
	"backend/backend/pkg/utils"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"golang.org/x/time/rate"
	"log"
	"time"
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

	limiter := rate.NewLimiter(rate.Every(time.Second*30), 3)
	webApp.Use(utils.MessageLimit(limiter))

	webApp.Post("/register", handlers.RegHandler(conn))
	webApp.Post("/entrance", handlers.EntranceHandler(conn))

	log.Fatal(webApp.Listen(":5174"))
}
