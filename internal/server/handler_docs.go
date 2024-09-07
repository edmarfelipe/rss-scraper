package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/edmarfelipe/rss-scraper/docs"
)

// DocsHandler is a handler for the API documentation.
func docsHandler(w http.ResponseWriter, r *http.Request) {
	content, err := docs.OpenAPISpec.ReadFile("api.yaml")
	if err != nil {
		slog.Error("failed to read OpenAPI spec", "error", err)
		http.Error(w, "failed to read OpenAPI spec", http.StatusInternalServerError)
		return
	}

	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecContent: string(content),
		DarkMode:    true,
	})
	if err != nil {
		slog.Error("failed to generate API reference HTML", "error", err)
		http.Error(w, "failed to generate API reference HTML", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, htmlContent)
}
