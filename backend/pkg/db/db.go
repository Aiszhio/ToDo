package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func InitDB() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), "postgres://postgres:1234@postgres:5432/usertodo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conn, nil
}
