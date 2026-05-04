package database

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Service interface {
}

type service struct {
	db *sql.DB
}

func New() Service {
	var db *sql.DB
	var err error
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		slog.Error("failed to connect to database", "err", "empty connection string")
		os.Exit(1)
	}
	if db, err = sql.Open("pgx", connStr); err != nil {
		slog.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}

	return &service{db: db}
}
