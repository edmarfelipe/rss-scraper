package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/edmarfelipe/rss-scraper/internal/database"
	"github.com/edmarfelipe/rss-scraper/internal/scraper"
	"github.com/edmarfelipe/rss-scraper/internal/server"
)

func main() {
	if err := run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run() error {
	db, err := database.NewConnection()
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}

	worker := scraper.NewScraper(db, 10, 60*time.Second)
	go worker.Start(context.Background())

	srv, err := server.NewServer(db)
	if err != nil {
		return fmt.Errorf("error creating server: %w", err)
	}

	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}
