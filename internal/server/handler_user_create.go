package server

import (
	"context"
	"time"

	"github.com/edmarfelipe/rss-scraper/internal/database"
	"github.com/edmarfelipe/rss-scraper/internal/server/openapi"
	"github.com/google/uuid"
)

var NameIsRequired = openapi.ErrorResponse{Error: "name is required"}

// UsersPost implements openapi.Handler.
func (s *Router) UsersPost(ctx context.Context, req *openapi.UsersPostReq) (openapi.UsersPostRes, error) {
	if req.Name == "" {
		return &openapi.UsersPostBadRequest{Error: "name is required"}, nil
	}

	user, err := s.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		Name:      req.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}

	return &openapi.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		APIKey:    user.ApiKey,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
