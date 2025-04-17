package main

import (
	"backend/backend/pkg/db"
	"backend/backend/pkg/handlers"
	"backend/backend/pkg/utils"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
	"log"
)

func main() {
	conn, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Ошибка подключения к Redis: %v", err)
	}
	fmt.Println("Подключение успешно:", pong)

	webApp := fiber.New()
	webApp.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE",
	}))

	limiter := rate.NewLimiter(1, 10)

	webApp.Post("/register", handlers.RegHandler(conn))
	webApp.Post("/entrance", utils.MessageLimit(limiter), handlers.EntranceHandler(conn, rdb, "secret"))
	webApp.Post("/notes", handlers.MakeNote(conn, "secret"))
	webApp.Post("/receive-token", handlers.ReceiveToken(rdb))
	webApp.Post("/user-notes", handlers.GetUserNotes(conn))

	defer conn.Close()
	defer func(rdb *redis.Client) {
		err = rdb.Close()
		if err != nil {

		}
	}(rdb)
	defer log.Fatal(webApp.Listen(":5174"))
}
