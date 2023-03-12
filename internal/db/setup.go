package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var dbURL = os.Getenv("DB_URL")

func New(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	return db, nil
}
