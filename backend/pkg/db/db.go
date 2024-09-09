package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func InitDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:1234@postgres:5432/usertodo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conn, nil
}
