package server

import (
	"context"

	"github.com/edmarfelipe/rss-scraper/internal/server/httputil"
	"github.com/edmarfelipe/rss-scraper/internal/server/openapi"
)

// FeedsFollowGet implements openapi.Handler.
func (s *Router) FeedsFollowGet(ctx context.Context) (openapi.FeedsFollowGetRes, error) {
	user := httputil.GetUser(ctx)
	if user == nil {
		return &openapi.FeedsFollowGetUnauthorized{Error: "user not authenticated"}, nil
	}

	feeds, err := s.DB.GetFeedByUser(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	response := openapi.FeedResponse{}
	for _, feed := range feeds {
		response.Feeds = append(response.Feeds, openapi.FeedResponseItem{
			ID:        feed.ID,
			Name:      feed.Name,
			URL:       feed.Url,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			UserID:    feed.UserID,
		})
	}

	return &response, nil
}
