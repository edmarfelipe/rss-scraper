package server

import (
	"context"

	"github.com/edmarfelipe/rss-scraper/internal/database"
	"github.com/edmarfelipe/rss-scraper/internal/server/httputil"
	"github.com/edmarfelipe/rss-scraper/internal/server/openapi"
	"github.com/google/uuid"
)

func (s *Router) FeedsFollowFeedFallowIDDelete(ctx context.Context, params openapi.FeedsFollowFeedFallowIDDeleteParams) (openapi.FeedsFollowFeedFallowIDDeleteRes, error) {
	if params.FeedFallowID == uuid.Nil {
		return &openapi.FeedsFollowFeedFallowIDDeleteBadRequest{Error: "feed_fallow_id is required"}, nil
	}

	user := httputil.GetUser(ctx)
	if user == nil {
		return &openapi.FeedsFollowFeedFallowIDDeleteBadRequest{Error: "user not authenticated"}, nil
	}

	err := s.DB.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		ID:     params.FeedFallowID,
		UserID: user.ID,
	})
	if err != nil {
		return nil, err
	}

	return &openapi.FeedsFollowFeedFallowIDDeleteNoContent{}, nil
}
