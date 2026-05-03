package main

import (
	"log/slog"
	"os"

	"github.com/alcb1310/bca-proto/internal/server"
)

func main() {
	var port = "8080"
	r, err := server.New(port)
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
