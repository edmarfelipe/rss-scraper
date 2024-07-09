package server

import (
	"context"
	"log/slog"

	"github.com/edmarfelipe/rss-scraper/internal/database"
	"github.com/edmarfelipe/rss-scraper/internal/server/httputil"
	"github.com/edmarfelipe/rss-scraper/internal/server/openapi"
)

// PostsGet implements openapi.Handler.
func (s *Router) PostsGet(ctx context.Context, params openapi.PostsGetParams) (openapi.PostsGetRes, error) {
	user := httputil.GetUser(ctx)
	if user == nil {
		return &openapi.PostsGetUnauthorized{Error: "user not authenticated"}, nil
	}

	posts, err := s.DB.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(params.Limit.Or(10)),
		Offset: int32(params.Page.Or(0) * params.Limit.Or(10)),
	})
	if err != nil {
		return nil, err
	}

	slog.Info("Got posts for user", "user_id", user.ID, "posts", len(posts))

	responseFeeds := []openapi.PostResponseItem{}
	for _, p := range posts {
		responseFeeds = append(responseFeeds, openapi.PostResponseItem{
			ID:          p.ID,
			Title:       p.Title,
			URL:         p.Url,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			PublishedAt: p.PublishedAt,
			FeedID:      p.FeedID,
			Content:     openapi.NewOptString(p.Content.String),
		})
	}

	return &openapi.PostResponse{Posts: responseFeeds}, nil
}
