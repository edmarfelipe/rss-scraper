package server

import (
	"context"
	"time"

	"github.com/edmarfelipe/rss-scraper/internal/database"
	"github.com/edmarfelipe/rss-scraper/internal/server/httputil"
	"github.com/edmarfelipe/rss-scraper/internal/server/openapi"
	"github.com/google/uuid"
)

func (s *Router) FeedsFollowPost(ctx context.Context, req *openapi.FeedsFollowPostReq) (openapi.FeedsFollowPostRes, error) {
	user := httputil.GetUser(ctx)
	if user == nil {
		return &openapi.FeedsFollowPostUnauthorized{Error: "user not authenticated"}, nil
	}

	if req.FeedID == "" {
		return &openapi.FeedsFollowPostBadRequest{Error: "feed_id is required"}, nil
	}

	feed, err := s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    uuid.MustParse(req.FeedID),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}

	return &openapi.FeedsFollowPostCreated{
		ID:        feed.ID,
		FeedID:    feed.FeedID,
		UserID:    feed.UserID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	}, nil
}
