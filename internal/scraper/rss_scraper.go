package scraper

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/edmarfelipe/rss-scraper/internal/database"
	"github.com/google/uuid"
)

type Scraper struct {
	db *database.Queries

	concurrency   int32
	fetchInterval time.Duration
}

func NewScraper(db *database.Queries, concurrency int32, fetchInterval time.Duration) *Scraper {
	return &Scraper{
		db:            db,
		concurrency:   concurrency,
		fetchInterval: fetchInterval,
	}
}

func (s *Scraper) Start(ctx context.Context) {
	slog.Info(fmt.Sprintf("Scraping on %d goroutines every %s", s.concurrency, s.fetchInterval))

	ticker := time.NewTicker(s.fetchInterval)
	for range ticker.C {
		feeds, err := s.db.GetNextFeedsToFetch(ctx, s.concurrency)
		if err != nil {
			slog.Error(fmt.Sprintf("error fetching feeds: %v", err))
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			slog.Info(fmt.Sprintf("Fetching feed %s", feed.Url))
			go func() {
				err := s.fetchFeed(ctx, wg, feed.ID, feed.Url)
				if err != nil {
					slog.Error(fmt.Sprintf("error fetching feed: %v", err))
				}
			}()

		}
		wg.Wait()
	}
}

func (s *Scraper) fetchFeed(ctx context.Context, wg *sync.WaitGroup, id uuid.UUID, url string) error {
	defer wg.Done()

	result, err := FetchFeed(url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	slog.Info(fmt.Sprintf("fetched %d items from feed %s", len(result.Channel.Items), url))

	err = s.db.MarkFeedAsFetched(ctx, id)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	for _, item := range result.Channel.Items {
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			slog.Error(fmt.Sprintf("error parsing date: %v", err))
			continue
		}

		_, err = s.db.CreatePost(ctx, database.CreatePostParams{
			ID:     uuid.New(),
			FeedID: id,
			Title:  item.Title,
			Url:    item.Link,
			Content: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: pubDate,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		})
		if err != nil && !database.IsUniqueViolation(err) {
			slog.Error(fmt.Sprintf("error creating item: %v", err))
			continue
		}
	}
	return nil
}
