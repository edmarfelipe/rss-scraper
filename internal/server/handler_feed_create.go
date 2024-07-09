package server

import (
	"context"

	"github.com/edmarfelipe/rss-scraper/internal/database"
	"github.com/edmarfelipe/rss-scraper/internal/server/httputil"
	"github.com/edmarfelipe/rss-scraper/internal/server/openapi"
	"github.com/google/uuid"
)

func (s *Router) FeedsPost(ctx context.Context, req *openapi.FeedsPostReq) (openapi.FeedsPostRes, error) {
	user := httputil.GetUser(ctx)
	if user == nil {
		return &openapi.FeedsPostUnauthorized{Error: "User not authenticated"}, nil
	}

	if req.Name == "" {
		return &openapi.FeedsPostBadRequest{Error: "name is required"}, nil
	}

	if req.URL == "" {
		return &openapi.FeedsPostBadRequest{Error: "url is required"}, nil
	}

	feed, err := s.DB.CreateFeed(ctx, database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   req.Name,
		Url:    req.URL,
		UserID: user.ID,
	})
	if err != nil {
		return nil, err
	}

	return &openapi.FeedResponseItem{
		ID:        feed.ID,
		Name:      feed.Name,
		URL:       feed.Url,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		UserID:    feed.UserID,
	}, nil
}
