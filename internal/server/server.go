package server

import (
	"net/http"
	"time"

	"github.com/edmarfelipe/rss-scraper/internal/database"
	"github.com/edmarfelipe/rss-scraper/internal/env"
	"github.com/edmarfelipe/rss-scraper/internal/server/mw"
)

func NewServer(db *database.Queries) (*http.Server, error) {
	mux, err := NewRouter(db)
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:         env.Config.ApiServerAddr,
		Handler:      mw.LoggingMiddleware(mux),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}, nil
}
