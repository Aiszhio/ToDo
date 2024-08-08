package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"log"
)

func regHandler(c *fiber.Ctx) error {
	if c.Method() == "POST" {
		email := c.FormValue("email")
		username := c.FormValue("username")
		password := c.FormValue("password")

		fmt.Println(email, username, password)
		return c.SendString("Register is OK" + email + "," + username + "," + password)
	}
	return c.SendString(`
        <form action="/register" method="post">
            <label for="email">Почта:</label>
            <input type="email" id="email" name="email" required><br><br>
            <label for="username">Имя пользователя:</label>
            <input type="text" id="username" name="username" required><br><br>
            <label for="password">Пароль:</label>
            <input type="password" id="password" name="password" required><br><br>
            <input type="submit" value="Зарегистрироваться">
        </form>
    `)
}

func main() {
	webApp := fiber.New()
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:-max656-@localhost:5432/user")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	webApp.Post("/", regHandler)
	webApp.Get("/", regHandler)
	log.Fatal(webApp.Listen(":5173"))
}
