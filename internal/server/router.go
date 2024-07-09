package server

import (
	"fmt"
	"net/http"

	"github.com/edmarfelipe/rss-scraper/internal/database"
	"github.com/edmarfelipe/rss-scraper/internal/server/openapi"
)

type Router struct {
	DB *database.Queries
}

func NewRouter(db *database.Queries) (*http.ServeMux, error) {
	srv, err := openapi.NewServer(&Router{DB: db}, newAuthHandler(db), openapi.WithErrorHandler(errorHandleFunc))
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", srv)
	mux.HandleFunc("GET /docs", docsHandler)
	return mux, nil
}
