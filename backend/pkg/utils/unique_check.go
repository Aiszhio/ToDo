package utils

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func IsUniqueEmail(email string, conn *pgx.Conn, ctx context.Context) bool {
	var exists bool = false
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`
	row := conn.QueryRow(ctx, query, email)
	if err := row.Scan(&exists); err != nil {
		return false
	}
	return exists
}

func IsUniqueName(name string, conn *pgx.Conn, ctx context.Context) bool {
	var exists bool = false
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`
	row := conn.QueryRow(ctx, query, name)
	if err := row.Scan(&exists); err != nil {
		return false
	}
	return exists
}
