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
	"time"
)

func main() {
	conn, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	rows, err := conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		log.Fatalf("failed to query users: %v", err)
	}
	fmt.Printf("Query: %+v\n", rows)

	for rows.Next() {
		var id int
		var email, name, password, hash_password string
		err = rows.Scan(&id, &email, &name, &password, &hash_password)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		fmt.Printf("User: %d, Email: %s, Name: %s, Password: %s, Hash_password: %s\n", id, email, name, password, hash_password)
	}
	rows.Close()

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
		AllowHeaders: "Origin, Content-Type",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE",
	}))

	limiter := rate.NewLimiter(rate.Every(time.Second*30), 3)
	webApp.Use(utils.MessageLimit(limiter))

	webApp.Post("/register", handlers.RegHandler(conn))
	webApp.Post("/entrance", handlers.EntranceHandler(conn, rdb, "secret"))
	webApp.Post("/notes", handlers.MakeNote(conn, "secret"))

	defer conn.Close()
	defer rdb.Close()
	defer log.Fatal(webApp.Listen(":5174"))
}
