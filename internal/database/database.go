package database

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Service interface {
	HealthCheck() error
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

func (s *service) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		if e, ok := errors.AsType[*pgconn.PgError](err); ok {
			return e
		}

		return err
	}

	return nil
}
