package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func InitDB() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), "postgres://postgres:1234@postgres:5432/usertodo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conn, nil
}

func ShowUsers(conn *pgxpool.Pool) error {
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

	return nil
}

func ShowNotes(conn *pgxpool.Pool) error {
	rows, err := conn.Query(context.Background(), "SELECT * FROM notes")
	if err != nil {
		log.Fatalf("failed to query notes: %v", err)
	}
	fmt.Printf("Query: %+v\n", rows)

	for rows.Next() {
		var id, user_id int
		var title string
		var created_to, created_at time.Time
		err = rows.Scan(&id, &user_id, &title, &created_to, &created_at)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		fmt.Printf("Id: %d, User_Id: %d, Title: %s, Created_to: %s, Created_at: %s\n", id, user_id, title, created_to, created_at)
	}
	rows.Close()

	return nil
}
