package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/edmarfelipe/rss-scraper/internal/database"
	"github.com/edmarfelipe/rss-scraper/internal/server/httputil"
	"github.com/edmarfelipe/rss-scraper/internal/server/openapi"
)

type AuthHandler struct {
	db *database.Queries
}

func newAuthHandler(db *database.Queries) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) HandleApiKeyAuth(ctx context.Context, operationName string, t openapi.ApiKeyAuth) (context.Context, error) {
	if t.GetAPIKey() == "" {
		return nil, fmt.Errorf("missing api key")
	}

	user, err := h.db.GetByAPIKey(ctx, t.GetAPIKey())
	if err != nil {
		return nil, fmt.Errorf("failed to get user by api key")
	}

	slog.Info(fmt.Sprintf("Authenticated user %s", user.Name))
	return httputil.WithUser(ctx, &user), nil
}
