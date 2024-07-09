package server

import (
	"context"

	"github.com/edmarfelipe/rss-scraper/internal/server/httputil"
	"github.com/edmarfelipe/rss-scraper/internal/server/openapi"
)

var UserNotAuthenticated = openapi.ErrorResponse{Error: "user not authenticated"}

// UsersGet implements openapi.Handler.
func (s *Router) UsersGet(ctx context.Context) (openapi.UsersGetRes, error) {
	user := httputil.GetUser(ctx)
	if user == nil {
		return &openapi.UsersGetNotFound{Error: "user not authenticated"}, nil
	}

	return &openapi.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
