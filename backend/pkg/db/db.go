package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func InitDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "user=postgres password=5589 host=localhost port=5432 dbname=postgres")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conn, nil
}
