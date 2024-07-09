package server

import (
	"context"

	"github.com/edmarfelipe/rss-scraper/internal/server/openapi"
)

// FeedsGet list all feeds.
func (s *Router) FeedsGet(ctx context.Context) (openapi.FeedsGetRes, error) {
	feeds, err := s.DB.GetFeeds(ctx)
	if err != nil {
		return nil, err
	}

	resp := openapi.FeedResponse{}
	for _, feed := range feeds {
		resp.Feeds = append(resp.Feeds, openapi.FeedResponseItem{
			ID:        feed.ID,
			Name:      feed.Name,
			URL:       feed.Url,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			UserID:    feed.UserID,
		})
	}
	return &resp, nil
}
