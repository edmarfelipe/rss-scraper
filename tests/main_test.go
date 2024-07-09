package tests

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/edmarfelipe/rss-scraper/internal/database"
	"github.com/edmarfelipe/rss-scraper/internal/server"
	"github.com/google/uuid"
)

var (
	srv *http.ServeMux
	db  *database.Queries
)

func TestMain(m *testing.M) {
	var err error
	db, err = database.NewConnection()
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %v", err))
	}

	srv, err = server.NewRouter(db)
	if err != nil {
		panic(fmt.Errorf("error creating server: %v", err))
	}

	exitVal := m.Run()

	teardown(context.Background())
	os.Exit(exitVal)
}

func teardown(ctx context.Context) {
	log.Println("tearing down database")

	db.TruncateFeeds(ctx)
	db.TruncateUsers(ctx)
	db.TruncatePosts(ctx)
}

func createUser(t *testing.T) database.User {
	t.Helper()

	r, err := db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      "Jon",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		t.Fatalf("error creating user: %v", err)
	}
	return r
}

func createFeed(t *testing.T) database.Feed {
	t.Helper()

	user := createUser(t)

	r, err := db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      "Feed Mock",
		Url:       "http://mock.com",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	})
	if err != nil {
		t.Fatalf("error creating feed: %v", err)
	}
	return r
}
