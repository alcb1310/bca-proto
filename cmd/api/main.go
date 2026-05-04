package main

import (
	"log/slog"
	"os"

	"github.com/alcb1310/bca-proto/internal/database"
	"github.com/alcb1310/bca-proto/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")

	db := database.New()

	r, err := server.New(port, db)
	if err != nil {
		slog.Error("failed to start server", "err", err)
		os.Exit(1)
	}

	slog.Info("server started", "port", port)
	if err := r.ListenAndServe(); err != nil {
		slog.Error("failed to start server", "err", err)
		os.Exit(1)
	}
}
