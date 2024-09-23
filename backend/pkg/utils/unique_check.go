package utils

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func IsUniqueEmail(email string, conn *pgxpool.Pool, ctx context.Context) bool {
	var exists bool = true
	row := conn.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`, email)
	if err := row.Scan(&exists); err != nil {
		return false
	}
	return exists
}

func IsUniqueName(name string, conn *pgxpool.Pool, ctx context.Context) bool {
	var exists bool = true
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE name = $1)`
	row := conn.QueryRow(ctx, query, name)
	if err := row.Scan(&exists); err != nil {
		return false
	}
	return exists
}
